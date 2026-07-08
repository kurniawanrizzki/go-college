package repository

import (
	"go-college/internal/infra/query"
	"go-college/internal/repository/college"
	"go-college/internal/repository/course"
	"go-college/internal/repository/enrollment"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type Repository struct {
	College    college.CollegeRepository
	Course     course.CourseRepository
	Enrollment enrollment.EnrollmentRepository
}

func InitRepository(log *zerolog.Logger, sql0 *pgxpool.Pool, queryLoader *query.QueryLoader) *Repository {
	repo := &Repository{
		College: college.InitCollegeRepository(
			log,
			sql0,
			queryLoader,
		),
		Course: course.InitCourseRepository(
			log,
			sql0,
			queryLoader,
		),
		Enrollment: enrollment.InitEnrollmentRepository(
			log,
			sql0,
			queryLoader,
		),
	}

	log.Info().Msg("repository initialized")

	return repo
}
