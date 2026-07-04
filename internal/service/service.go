package service

import (
	"go-college/internal/repository"
	"go-college/internal/service/college"
	"net/http"

	"github.com/rs/zerolog"
)

type Service struct {
	College college.CollegeService
}

func InitService(log *zerolog.Logger, repo *repository.Repository, httpClient *http.Client) *Service {
	svc := &Service{
		College: college.InitCollegeService(
			log,
			repo.College,
		),
	}

	log.Info().Msg("service initialized")

	return svc
}
