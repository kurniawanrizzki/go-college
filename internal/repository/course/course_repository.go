package course

import (
	"context"

	"go-college/internal/model/dto"
	"go-college/internal/model/entity"
	appErr "go-college/internal/model/errors"
	"go-college/internal/util"

	"github.com/rs/zerolog"
)

func (d *courseRepositoryImpl) Create(ctx context.Context, course *entity.Course) (*entity.Course, error) {
	tx, err := d.sql0.Begin(ctx)
	if err != nil {
		zerolog.Ctx(ctx).Error().Msg("tx_create_course")
		return course, appErr.Wrap(err, "tx_create_course")
	}

	rolledBack := false

	defer func() {
		if !rolledBack {
			_ = tx.Rollback(ctx)
		}
	}()

	tx, course, err = d.createSQLCourse(ctx, tx, course)

	if err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("sql_create_course")
		return nil, appErr.Wrap(err, "sql_create_course")
	}

	if err = tx.Commit(ctx); err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("commit_create_course")
		return nil, appErr.Wrap(err, "commit_create_course")
	}

	rolledBack = true

	return course, nil
}

func (d *courseRepositoryImpl) FindByCode(ctx context.Context, code string) (*entity.Course, error) {
	return d.findSQLCourseByCode(ctx, code)
}

func (d *courseRepositoryImpl) Update(ctx context.Context, course *entity.Course) error {
	query, args, err := d.queryLoader.Compile("UpdateCourse", course)
	return d.executeSQLCourse(ctx, "update", query, args, err)
}

func (d *courseRepositoryImpl) Delete(ctx context.Context, code string) error {
	query, args, err := d.queryLoader.Compile("DeleteCourse", map[string]any{"Code": code})
	return d.executeSQLCourse(ctx, "delete", query, args, err)
}

func (d *courseRepositoryImpl) FindAll(ctx context.Context, filter *dto.CourseFilter) (*[]entity.Course, *dto.Pagination, error) {
	var (
		results      []entity.Course
		totalRecords int64
	)

	filter.Page = util.ValidatePage(filter.Page)
	filter.PerPage = util.ValidateLimit(filter.PerPage)
	filter.SortBy = sanitizeSortBy(filter.SortBy)
	filter.SortDir = sanitizeSortDir(filter.SortDir)
	filter.Limit = filter.PerPage
	filter.Offset = (filter.Page - 1) * filter.PerPage

	pagination := dto.Pagination{
		Page:       filter.Page,
		PageCount:  0,
		TotalCount: 0,
		PerPage:    filter.PerPage,
	}

	query, args, err := d.queryLoader.Compile("GetAllCourses", filter)

	if err != nil {
		tag := "build_find_courses_err"
		zerolog.Ctx(ctx).Error().Err(err).Msg(tag)
		return nil, &pagination, appErr.WrapWithCode(err, appErr.CodeSQLQueryBuild, tag)
	}

	zerolog.Ctx(ctx).Debug().Str("query", util.CleanQuery(query)).Any("args", args).Msg("compiled_query")

	rows, err := d.sql0.Query(ctx, query, args...)

	if err != nil {
		tag := "find_courses_err"
		zerolog.Ctx(ctx).Error().Err(err).Msg(tag)
		return nil, &pagination, appErr.WrapWithCode(err, appErr.CodeSQLRowScan, tag)
	}
	defer rows.Close()

	for rows.Next() {
		var course entity.Course
		if scanErr := rows.Scan(&course.Code, &course.Name, &course.SKS, &course.CreatedAt, &course.UpdatedAt); scanErr != nil {
			zerolog.Ctx(ctx).Error().Err(scanErr).Msg("scan_course_err")
			return nil, nil, appErr.WrapWithCode(scanErr, appErr.CodeSQLRowScan, "scan_course_err")
		}

		results = append(results, course)
	}

	cQuery, cArgs, err := d.queryLoader.Compile("CountCourses", filter)

	if err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("count_courses_query_err")
		return nil, &pagination, appErr.WrapWithCode(err, appErr.CodeSQLQueryBuild, "count_courses_query_err")
	}

	err = d.sql0.QueryRow(ctx, cQuery, cArgs...).Scan(&totalRecords)
	if err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("count_courses_err")
		return nil, &pagination, appErr.WrapWithCode(err, appErr.CodeSQLRowScan, "count_courses_err")
	}

	var totalPage int64
	if totalRecords > 0 {
		totalPage = (totalRecords + filter.PerPage - 1) / filter.PerPage
	} else {
		totalPage = 0
	}

	pagination.PageCount = totalPage
	pagination.TotalCount = totalRecords

	return &results, &pagination, nil
}
