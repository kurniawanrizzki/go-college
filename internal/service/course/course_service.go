package course

import (
	"context"

	"go-college/internal/model/dto"
	"go-college/internal/model/entity"
	appErr "go-college/internal/model/errors"
)

func (s *courseServiceImpl) Create(ctx context.Context, req dto.CreateCourseRequest) (*entity.Course, error) {
	course, err := s.repository.Create(
		ctx,
		&entity.Course{
			Code: req.Code,
			Name: req.Name,
			SKS:  req.SKS,
		},
	)

	if err != nil {
		return nil, err
	}

	return course, nil
}

func (s *courseServiceImpl) FindByCode(ctx context.Context, code string) (*entity.Course, error) {
	return s.repository.FindByCode(ctx, code)
}

func (s *courseServiceImpl) Update(ctx context.Context, code string, req *dto.UpdateCourseRequest) (*entity.Course, error) {
	existing, err := s.FindByCode(ctx, code)

	if err != nil {
		return nil, appErr.WrapWithCode(err, appErr.CodeHTTPNotFound, "college_not_found")
	}

	if req.Name != "" {
		existing.Name = req.Name
	}

	if req.SKS != 0 {
		existing.SKS = req.SKS
	}

	if err = s.repository.Update(
		ctx,
		existing,
	); err != nil {
		return nil, err
	}

	return existing, nil
}

func (s *courseServiceImpl) Delete(ctx context.Context, code string) error {
	return s.repository.Delete(ctx, code)
}

func (s *courseServiceImpl) FindAll(ctx context.Context) (*[]entity.Course, error) {
	return s.repository.FindAll(ctx)
}
