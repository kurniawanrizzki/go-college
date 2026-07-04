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
