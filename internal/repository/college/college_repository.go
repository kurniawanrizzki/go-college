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
	return d.findSQLByArgs(ctx, "colleges", "FindColleges", nil)
}

func (d *collegeRepositoryImpl) Update(ctx context.Context, college *entity.College) error {
	query, args, err := d.queryLoader.Compile("UpdateCollege", college)
	return d.executeSQLCollege(ctx, "update", query, args, err)
}

func (d *collegeRepositoryImpl) Delete(ctx context.Context, nim string) error {
	query, args, err := d.queryLoader.Compile("DeleteCollege", map[string]any{"NIM": nim})
	return d.executeSQLCollege(ctx, "delete", query, args, err)
}

func (d *collegeRepositoryImpl) FindByNim(ctx context.Context, nim string) (*entity.College, error) {
	return d.findSQLCollegeByNIM(ctx, nim)
}

func (d *collegeRepositoryImpl) FindByName(ctx context.Context, name string) (*[]entity.College, error) {
	return d.findSQLByArgs(ctx, "nim", "FindCollegeByName", map[string]any{"Name": name})
}

func (d *collegeRepositoryImpl) FindBySemester(ctx context.Context, semester int) (*[]entity.College, error) {
	return d.findSQLByArgs(ctx, "semester", "FindCollegeBySemester", map[string]any{"Semester": semester})
}
