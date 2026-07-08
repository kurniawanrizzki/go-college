package enrollment

import (
	"context"
	"fmt"
	"go-college/internal/model/entity"
	appErr "go-college/internal/model/errors"
	"go-college/internal/util"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
)

func (d *enrollmentRepositoryImpl) createSQLEnrollment(ctx context.Context, tx pgx.Tx, enrollment *entity.Enrollment) (pgx.Tx, *entity.Enrollment, error) {
	query, args, err := d.queryLoader.Compile("CreateEnrollment", enrollment)

	if err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("build_create_enrollment_query_err")
		return tx, enrollment, appErr.WrapWithCode(err, appErr.CodeSQLQueryBuild, "build_create_enrollment_query_err")
	}

	zerolog.Ctx(ctx).Debug().Str("query", util.CleanQuery(query)).Any("args", args).Msg("compiled_query")

	err = tx.QueryRow(ctx, query, args...).Scan(&enrollment.ID, &enrollment.Grade, &enrollment.CreatedAt, &enrollment.UpdatedAt)

	if err != nil {
		return tx, enrollment, appErr.Wrap(err, "create_sql_enrollment")
	}

	return tx, enrollment, nil
}

func (d *enrollmentRepositoryImpl) executeSQLEnrollment(ctx context.Context, key string, query string, args []any, err error) error {
	if err != nil {
		tag := fmt.Sprintf("build_%v_enrollment", key)
		zerolog.Ctx(ctx).Error().Err(err).Msg(tag)
		return appErr.WrapWithCode(err, appErr.CodeSQLBuilder, tag)
	}

	zerolog.Ctx(ctx).Debug().Str("query", util.CleanQuery(query)).Any("args", args).Msg("compiled_query")

	result, _ := d.sql0.Exec(ctx, query, args...)
	rows := result.RowsAffected()

	if rows == 0 {
		tag := fmt.Sprintf("enrollment_not_found_for_%v", key)
		return appErr.NewWithCode(appErr.CodeSQLEmptyRow, tag)
	}

	return nil
}

func (d *enrollmentRepositoryImpl) findSQLDetailByNim(ctx context.Context, nim string) (*[]entity.EnrollmentDetail, error) {
	var results []entity.EnrollmentDetail

	query, args, err := d.queryLoader.Compile("FindEnrollmentDetailByNim", map[string]any{"NIM": nim})

	if err != nil {
		tag := "build_find_enrollment_detail_err"
		zerolog.Ctx(ctx).Error().Err(err).Msg(tag)
		return nil, appErr.WrapWithCode(err, appErr.CodeSQLQueryBuild, tag)
	}

	zerolog.Ctx(ctx).Debug().Str("query", util.CleanQuery(query)).Any("args", args).Msg("compiled_query")

	rows, err := d.sql0.Query(ctx, query, args...)

	if err != nil {
		tag := "find_enrollment_detail_err"
		zerolog.Ctx(ctx).Error().Err(err).Msg(tag)
		return nil, appErr.WrapWithCode(err, appErr.CodeSQLRowScan, tag)
	}
	defer rows.Close()

	for rows.Next() {
		var detail entity.EnrollmentDetail
		if scanErr := rows.Scan(
			&detail.Course.Code,
			&detail.Course.Name,
			&detail.Course.SKS,
			&detail.Course.CreatedAt,
			&detail.Course.UpdatedAt,
			&detail.Semester,
			&detail.Grade,
			&detail.CreatedAt,
			&detail.UpdatedAt,
		); scanErr != nil {
			zerolog.Ctx(ctx).Error().Err(scanErr).Msg("scan_enrollment_detail_err")
			return nil, appErr.WrapWithCode(scanErr, appErr.CodeSQLRowScan, "scan_enrollment_detail_err")
		}

		results = append(results, detail)
	}

	return &results, nil
}
