package college

import (
	"context"

	"go-college/internal/model/dto"
	"go-college/internal/model/entity"
	appErr "go-college/internal/model/errors"
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

func (s *collegeServiceImpl) FindAll(ctx context.Context, filter *dto.CollegeFilter) (*[]entity.College, *dto.Pagination, error) {
	return s.repository.FindAll(ctx, filter)
}

func (s *collegeServiceImpl) Update(ctx context.Context, nim string, req *dto.UpdateCollegeRequest) (*entity.College, error) {
	existing, err := s.repository.FindByNim(ctx, nim)

	if err != nil {
		return nil, appErr.WrapWithCode(err, appErr.CodeHTTPNotFound, "college_not_found")
	}

	if req.Name != "" {
		existing.Name = req.Name
	}

	if req.SKS != 0 {
		existing.SKS = req.SKS
	}

	if req.Semester != 0 {
		existing.Semester = req.Semester
	}

	if req.Active != nil {
		existing.Active = *req.Active
	}

	if err = s.repository.Update(
		ctx,
		existing,
	); err != nil {
		return nil, err
	}

	return existing, nil
}

func (s *collegeServiceImpl) Delete(ctx context.Context, nim string) error {
	return s.repository.Delete(ctx, nim)
}
