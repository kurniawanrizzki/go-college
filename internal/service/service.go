package service

import (
	"go-college/internal/repository"
	"go-college/internal/service/college"
	"go-college/internal/service/course"
	"net/http"

	"github.com/rs/zerolog"
)

type Service struct {
	College college.CollegeService
	Course  course.CourseService
}

func InitService(log *zerolog.Logger, repo *repository.Repository, httpClient *http.Client) *Service {
	svc := &Service{
		College: college.InitCollegeService(
			log,
			repo.College,
		),
		Course: course.InitCourseService(
			log,
			repo.Course,
		),
	}

	log.Info().Msg("service initialized")

	return svc
}
