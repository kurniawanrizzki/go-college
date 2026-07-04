package repository

import (
	"go-college/internal/infra/query"
	"go-college/internal/repository/college"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type Repository struct {
	College college.CollegeRepository
}

func InitRepository(log *zerolog.Logger, sql0 *pgxpool.Pool, queryLoader *query.QueryLoader) *Repository {
	repo := &Repository{
		College: college.InitCollegeRepository(
			log,
			sql0,
			queryLoader,
		),
	}

	log.Info().Msg("repository initialized")

	return repo
}
