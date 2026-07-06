package college

import (
	"context"
	"go-college/internal/model/dto"
	"go-college/internal/model/entity"
)

func (s *collegeServiceImpl) Create(ctx context.Context, req dto.CreateCollegeRequest) (*entity.College, error) {
	college, err := s.repository.Create(
		ctx,
		&entity.College{
			NIM:      req.NIM,
			Name:     req.Name,
			Semester: req.Semester,
			SKS:      req.SKS,
			Active:   req.Active,
		},
	)

	if err != nil {
		return nil, err
	}

	return college, nil
}

func (s *collegeServiceImpl) FindAll(ctx context.Context) (*[]entity.College, error) {
	return s.repository.FindAll(ctx)
}
