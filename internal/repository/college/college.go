package college

import (
	"context"

	"go-college/internal/infra/query"
	"go-college/internal/model/entity"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type CollegeRepository interface {
	Create(ctx context.Context, college *entity.College) (*entity.College, error)
	Update(ctx context.Context, college *entity.College) error
	Delete(ctx context.Context, nim string) error
	FindByNim(ctx context.Context, nim string) (*entity.College, error)
	FindByName(ctx context.Context, name string) (*[]entity.College, error)
	FindBySemester(ctx context.Context, semester int) (*[]entity.College, error)
	FindAll(ctx context.Context) (*[]entity.College, error)
}

type collegeRepositoryImpl struct {
	log         *zerolog.Logger
	sql0        *pgxpool.Pool
	queryLoader *query.QueryLoader
}

func InitCollegeRepository(log *zerolog.Logger, sql0 *pgxpool.Pool, queryLoader *query.QueryLoader) CollegeRepository {
	return &collegeRepositoryImpl{
		log:         log,
		sql0:        sql0,
		queryLoader: queryLoader,
	}
}
