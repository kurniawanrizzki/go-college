package service

import (
	"net/http"

	"go-college/internal/repository"
	"go-college/internal/service/college"
	"go-college/internal/service/course"
	"go-college/internal/service/enrollment"

	"github.com/rs/zerolog"
)

type Service struct {
	College    college.CollegeService
	Course     course.CourseService
	Enrollment enrollment.EnrollmentService
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
		Enrollment: enrollment.InitEnrollmentService(
			log,
			repo.Enrollment,
		),
	}

	log.Info().Msg("service initialized")

	return svc
}
