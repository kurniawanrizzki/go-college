package enrollment

import (
	"context"

	"go-college/internal/model/dto"
	"go-college/internal/model/entity"
	"go-college/internal/repository/enrollment"

	"github.com/rs/zerolog"
)

type EnrollmentService interface {
	Create(ctx context.Context, req dto.CreateEnrollmentRequest) (*entity.Enrollment, error)
	Update(ctx context.Context, nim string, course string, req *dto.UpdateEnrollmentRequest) error
	Delete(ctx context.Context, id int) error
	FindDetailByNim(ctx context.Context, nim string) (*[]entity.EnrollmentDetail, error)
}

type enrollmentServiceImpl struct {
	log        *zerolog.Logger
	repository enrollment.EnrollmentRepository
}

func InitEnrollmentService(log *zerolog.Logger, repository enrollment.EnrollmentRepository) EnrollmentService {
	return &enrollmentServiceImpl{
		log:        log,
		repository: repository,
	}
}
