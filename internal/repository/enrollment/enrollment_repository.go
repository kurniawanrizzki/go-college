package enrollment

import (
	"context"
	"go-college/internal/model/entity"
	appErr "go-college/internal/model/errors"

	"github.com/rs/zerolog"
)

func (d *enrollmentRepositoryImpl) Create(ctx context.Context, enrollment *entity.Enrollment) (*entity.Enrollment, error) {
	tx, err := d.sql0.Begin(ctx)
	if err != nil {
		zerolog.Ctx(ctx).Error().Msg("tx_create_enrollment")
		return enrollment, appErr.Wrap(err, "tx_create_enrollment")
	}

	rolledBack := false

	defer func() {
		if !rolledBack {
			_ = tx.Rollback(ctx)
		}
	}()

	tx, enrollment, err = d.createSQLEnrollment(ctx, tx, enrollment)

	if err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("sql_create_enrollment")
		return nil, appErr.Wrap(err, "sql_create_enrollment")
	}

	if err = tx.Commit(ctx); err != nil {
		zerolog.Ctx(ctx).Error().Err(err).Msg("commit_create_enrollment")
		return nil, appErr.Wrap(err, "commit_create_enrollment")
	}

	rolledBack = true

	return enrollment, nil
}

func (d *enrollmentRepositoryImpl) Update(ctx context.Context, enrollment *entity.Enrollment) error {
	query, args, err := d.queryLoader.Compile("UpdateEnrollment", enrollment)
	return d.executeSQLEnrollment(ctx, "update", query, args, err)
}

func (d *enrollmentRepositoryImpl) Delete(ctx context.Context, id int) error {
	query, args, err := d.queryLoader.Compile("DeleteEnrollment", map[string]any{"ID": id})
	return d.executeSQLEnrollment(ctx, "delete", query, args, err)
}

func (d *enrollmentRepositoryImpl) FindDetailByNim(ctx context.Context, nim string) (*[]entity.EnrollmentDetail, error) {
	return d.findSQLDetailByNim(ctx, nim)
}
