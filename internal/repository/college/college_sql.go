package college

import (
	"context"
	"errors"
	"fmt"
	"go-college/internal/model/entity"
	appErr "go-college/internal/model/errors"
	"go-college/internal/util"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
)

func (d *collegeRepositoryImpl) createSQLCollege(ctx context.Context, tx pgx.Tx, college *entity.College) (pgx.Tx, *entity.College, error) {
	query, args, err := d.queryLoader.Compile("CreateCollege", college)

	if err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("build_create_college_query_err")
		return tx, college, appErr.WrapWithCode(err, appErr.CodeSQLQueryBuild, "build_create_college_query_err")
	}

	zerolog.Ctx(ctx).Debug().Str("query", util.CleanQuery(query)).Any("args", args).Msg("compiled_query")

	err = tx.QueryRow(ctx, query, args...).Scan(&college.CreatedAt, &college.UpdatedAt)

	if err != nil {
		return tx, college, appErr.Wrap(err, "create_sql_college")
	}

	return tx, college, nil
}

func (d *collegeRepositoryImpl) findSQLCollegeByNIM(ctx context.Context, nim string) (*entity.College, error) {
	var college entity.College

	query, args, err := d.queryLoader.Compile("FindCollegeByNim", map[string]any{"NIM": nim})
	if err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("build_find_college_by_nim_query_err")
		return nil, appErr.WrapWithCode(err, appErr.CodeSQLQueryBuild, "build_find_college_by_nim_query_err")
	}

	zerolog.Ctx(ctx).Debug().Str("query", util.CleanQuery(query)).Any("args", args).Msg("compiled_query")

	err = d.sql0.QueryRow(ctx, query, args...).Scan(
		&college.NIM,
		&college.Name,
		&college.Semester,
		&college.SKS,
		&college.Active,
		&college.CreatedAt,
		&college.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, appErr.NewWithCode(appErr.CodeHTTPUnauthorized, "invalid credentials")
		}

		zerolog.Ctx(ctx).Error().Err(err).Str("nim", nim).Msg("find_college_by_nim_err")
		return nil, appErr.WrapWithCode(err, appErr.CodeSQLRowScan, "find_college_by_nim_err")
	}

	return &college, nil
}

func (d *collegeRepositoryImpl) findSQLByArgs(ctx context.Context, key string, query string, data map[string]any) (*[]entity.College, error) {
	var results []entity.College

	query, args, err := d.queryLoader.Compile(query, data)

	if err != nil {
		tag := fmt.Sprintf("build_find_%v_err", key)
		zerolog.Ctx(ctx).Error().Err(err).Msg(tag)
		return nil, appErr.WrapWithCode(err, appErr.CodeSQLQueryBuild, tag)
	}

	zerolog.Ctx(ctx).Debug().Str("query", util.CleanQuery(query)).Any("args", args).Msg("compiled_query")

	rows, err := d.sql0.Query(ctx, query, args...)

	if err != nil {
		tag := fmt.Sprintf("find_%v_err", key)
		zerolog.Ctx(ctx).Error().Err(err).Msg(tag)
		return nil, appErr.WrapWithCode(err, appErr.CodeSQLRowScan, tag)
	}
	defer rows.Close()

	for rows.Next() {
		var college entity.College
		if scanErr := rows.Scan(&college.NIM, &college.Name, &college.Semester, &college.SKS, &college.Active, &college.CreatedAt, &college.UpdatedAt); scanErr != nil {
			zerolog.Ctx(ctx).Error().Err(scanErr).Msg("scan_college_err")
			return nil, appErr.WrapWithCode(scanErr, appErr.CodeSQLRowScan, "scan_college_err")
		}

		results = append(results, college)
	}

	return &results, nil
}

func (d *collegeRepositoryImpl) executeSQLCollege(ctx context.Context, key string, query string, args []any, err error) error {
	if err != nil {
		tag := fmt.Sprintf("build_%v_college", key)
		zerolog.Ctx(ctx).Error().Err(err).Msg(tag)
		return appErr.WrapWithCode(err, appErr.CodeSQLBuilder, tag)
	}

	zerolog.Ctx(ctx).Debug().Str("query", util.CleanQuery(query)).Any("args", args).Msg("compiled_query")

	result, _ := d.sql0.Exec(ctx, query, args...)
	rows := result.RowsAffected()

	if rows == 0 {
		tag := fmt.Sprintf("college_not_found_for_%v", key)
		return appErr.NewWithCode(appErr.CodeSQLEmptyRow, tag)
	}

	return nil
}
