package college

import (
	"context"
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

func (d *collegeRepositoryImpl) findSQLColleges(ctx context.Context) (*[]entity.College, error) {
	var results []entity.College

	query, args, err := d.queryLoader.Compile("FindColleges", nil)

	if err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("build_find_colleges_err")
		return nil, appErr.WrapWithCode(err, appErr.CodeSQLQueryBuild, "build_find_colleges_err")
	}

	zerolog.Ctx(ctx).Debug().Str("query", util.CleanQuery(query)).Any("args", args).Msg("compiled_query")

	rows, err := d.sql0.Query(ctx, query, args...)

	if err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("find_colleges_err")
		return nil, appErr.WrapWithCode(err, appErr.CodeSQLRowScan, "find_colleges_err")
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
