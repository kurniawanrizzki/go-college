package course

import (
	"context"

	"go-college/internal/model/dto"
	"go-college/internal/model/entity"
	"go-college/internal/repository/course"

	"github.com/rs/zerolog"
)

type CourseService interface {
	Create(ctx context.Context, req dto.CreateCourseRequest) (*entity.Course, error)
	Update(ctx context.Context, code string, req *dto.UpdateCourseRequest) (*entity.Course, error)
	Delete(ctx context.Context, code string) error
	FindAll(ctx context.Context, filter *dto.CourseFilter) (*[]entity.Course, *dto.Pagination, error)
}

type courseServiceImpl struct {
	log        *zerolog.Logger
	repository course.CourseRepository
}

func InitCourseService(log *zerolog.Logger, repository course.CourseRepository) CourseService {
	return &courseServiceImpl{
		log:        log,
		repository: repository,
	}
}
