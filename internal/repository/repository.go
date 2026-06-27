package repository

import (
	"go-college/internal/infra/query"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type Repository struct {}

func InitRepository(log *zerolog.Logger, sql0 *pgxpool.Pool, queryLoader *query.QueryLoader) *Repository {
	repo := &Repository{}

	log.Info().Msg("repository initialized")

	return repo
}
