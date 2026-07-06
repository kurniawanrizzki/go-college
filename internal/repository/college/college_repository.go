package college

import (
	"context"
	"go-college/internal/model/entity"
	appErr "go-college/internal/model/errors"

	"github.com/rs/zerolog"
)

func (d *collegeRepositoryImpl) Create(ctx context.Context, college *entity.College) (*entity.College, error) {
	tx, err := d.sql0.Begin(ctx)
	if err != nil {
		zerolog.Ctx(ctx).Error().Msg("tx_create_college")
		return college, appErr.Wrap(err, "tx_create_college")
	}
	rolledBack := false
	
	defer func() {
		if !rolledBack {
			_ = tx.Rollback(ctx)
		}
	}()

	tx, college, err = d.createSQLCollege(ctx, tx, college)

	if err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("sql_create_college")
		return nil, appErr.Wrap(err, "sql_create_college")
	}

	if err = tx.Commit(ctx); err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("commit_create_college")
		return nil, appErr.Wrap(err, "commit_create_college")
	}

	rolledBack = true

	return college, nil
}

func (d *collegeRepositoryImpl) FindAll(ctx context.Context) (*[]entity.College, error) {
	return d.findSQLColleges(ctx)
}
