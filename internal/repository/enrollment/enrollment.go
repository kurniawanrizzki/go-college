package enrollment

import (
	"context"

	"go-college/internal/infra/query"
	"go-college/internal/model/entity"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type EnrollmentRepository interface {
	Create(ctx context.Context, enrollment *entity.Enrollment) (*entity.Enrollment, error)
	Update(ctx context.Context, enrollment *entity.Enrollment) error
	Delete(ctx context.Context, id int) error
	FindDetailByNim(ctx context.Context, nim string) (*[]entity.EnrollmentDetail, error)
}

type enrollmentRepositoryImpl struct {
	log         *zerolog.Logger
	sql0        *pgxpool.Pool
	queryLoader *query.QueryLoader
}

func InitEnrollmentRepository(log *zerolog.Logger, sql0 *pgxpool.Pool, queryLoader *query.QueryLoader) EnrollmentRepository {
	return &enrollmentRepositoryImpl{
		log:         log,
		sql0:        sql0,
		queryLoader: queryLoader,
	}
}
