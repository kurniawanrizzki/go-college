package course

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"go-college/internal/model/entity"
	appErr "go-college/internal/model/errors"
	"go-college/internal/util"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
)

var (
	allowedSortFields = map[string]string{
		"code": "code",
		"name": "name",
		"sks":  "sks",
	}

	allowedSortDirs = map[string]string{
		"asc":  "ASC",
		"desc": "DESC",
	}
)

func sanitizeSortBy(sortBy string) string {
	normalized := normalizeString(sortBy)
	if col, ok := allowedSortFields[normalized]; ok {
		return col
	}

	return "code"
}

func sanitizeSortDir(sortDir string) string {
	normalized := normalizeString(sortDir)
	if dir, ok := allowedSortDirs[normalized]; ok {
		return dir
	}

	return "ASC"
}

func normalizeString(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	return s
}

func (d *courseRepositoryImpl) createSQLCourse(context context.Context, tx pgx.Tx, course *entity.Course) (pgx.Tx, *entity.Course, error) {
	query, args, err := d.queryLoader.Compile("CreateCourse", course)

	if err != nil {
		zerolog.Ctx(context).Error().Err(err).Msg("build_create_course_query_err")
		return tx, course, appErr.WrapWithCode(err, appErr.CodeSQLQueryBuild, "build_create_course_query_err")
	}

	zerolog.Ctx(context).Debug().Str("query", util.CleanQuery(query)).Any("args", args).Msg("compiled_query")

	err = tx.QueryRow(context, query, args...).Scan(&course.CreatedAt, &course.UpdatedAt)

	if err != nil {
		return tx, course, appErr.Wrap(err, "create_sql_course")
	}

	return tx, course, nil
}

func (d *courseRepositoryImpl) executeSQLCourse(ctx context.Context, key string, query string, args []any, err error) error {
	if err != nil {
		tag := fmt.Sprintf("build_%v_course", key)
		zerolog.Ctx(ctx).Error().Err(err).Msg(tag)
		return appErr.WrapWithCode(err, appErr.CodeSQLBuilder, tag)
	}

	zerolog.Ctx(ctx).Debug().Str("query", util.CleanQuery(query)).Any("args", args).Msg("compiled_query")

	result, _ := d.sql0.Exec(ctx, query, args...)
	rows := result.RowsAffected()

	if rows == 0 {
		tag := fmt.Sprintf("course_not_found_for_%v", key)
		return appErr.NewWithCode(appErr.CodeSQLEmptyRow, tag)
	}

	return nil
}

func (d *courseRepositoryImpl) findSQLCourseByCode(ctx context.Context, code string) (*entity.Course, error) {
	var course entity.Course

	query, args, err := d.queryLoader.Compile("FindCourseByCode", map[string]any{"Code": code})
	if err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("build_find_course_by_code_query_err")
		return nil, appErr.WrapWithCode(err, appErr.CodeSQLQueryBuild, "build_find_course_by_code_query_err")
	}

	zerolog.Ctx(ctx).Debug().Str("query", util.CleanQuery(query)).Any("args", args).Msg("compiled_query")

	err = d.sql0.QueryRow(ctx, query, args...).Scan(
		&course.Code,
		&course.Name,
		&course.SKS,
		&course.CreatedAt,
		&course.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, appErr.NewWithCode(appErr.CodeHTTPUnauthorized, "invalid credentials")
		}

		zerolog.Ctx(ctx).Error().Err(err).Str("code", code).Msg("find_course_by_code_err")
		return nil, appErr.WrapWithCode(err, appErr.CodeSQLRowScan, "find_course_by_code_err")
	}

	return &course, nil
}
