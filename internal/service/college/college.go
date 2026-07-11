package college

import (
	"context"

	"go-college/internal/model/dto"
	"go-college/internal/model/entity"
	"go-college/internal/repository/college"

	"github.com/rs/zerolog"
)

type CollegeService interface {
	Create(ctx context.Context, req dto.CreateCollegeRequest) (*entity.College, error)
	FindAll(ctx context.Context, filter *dto.CollegeFilter) (*[]entity.College, *dto.Pagination, error)
	Update(ctx context.Context, nim string, req *dto.UpdateCollegeRequest) (*entity.College, error)
	Delete(ctx context.Context, nim string) error
}

type collegeServiceImpl struct {
	log        *zerolog.Logger
	repository college.CollegeRepository
}

func InitCollegeService(log *zerolog.Logger, repository college.CollegeRepository) CollegeService {
	return &collegeServiceImpl{
		log:        log,
		repository: repository,
	}
}
