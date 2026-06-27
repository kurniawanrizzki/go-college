package service

import (
	"go-college/internal/repository"
	"net/http"

	"github.com/rs/zerolog"
)

type Service struct {}

func InitService(log *zerolog.Logger, repo *repository.Repository, httpClient *http.Client) *Service {
	svc := &Service{}

	log.Info().Msg("service initialized")

	return svc
}
