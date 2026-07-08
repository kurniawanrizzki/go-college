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
	FindAll(ctx context.Context) (*[]entity.College, error)
	Update(ctx context.Context, nim string, req *dto.UpdateCollegeRequest) (*entity.College, error)
	Delete(ctx context.Context, nim string) error
	FindByNim(ctx context.Context, nim string) (*entity.College, error)
	FindByName(ctx context.Context, name string) (*[]entity.College, error)
	FindBySemester(ctx context.Context, semester int) (*[]entity.College, error)
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
