package course

import (
	"context"

	"go-college/internal/infra/query"
	"go-college/internal/model/entity"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type CourseRepository interface {
	Create(ctx context.Context, course *entity.Course) (*entity.Course, error)
	Update(ctx context.Context, course *entity.Course) error
	Delete(ctx context.Context, code string) error
	FindAll(ctx context.Context) (*[]entity.Course, error)
	FindByCode(ctx context.Context, code string) (*entity.Course, error)
}

type courseRepositoryImpl struct {
	log         *zerolog.Logger
	sql0        *pgxpool.Pool
	queryLoader *query.QueryLoader
}

func InitCourseRepository(log *zerolog.Logger, sql0 *pgxpool.Pool, queryLoader *query.QueryLoader) CourseRepository {
	return &courseRepositoryImpl{
		log:         log,
		sql0:        sql0,
		queryLoader: queryLoader,
	}
}
