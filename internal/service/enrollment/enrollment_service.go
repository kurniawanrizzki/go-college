package enrollment

import (
	"context"

	"go-college/internal/model/dto"
	"go-college/internal/model/entity"
)

func (s *enrollmentServiceImpl) Create(ctx context.Context, req dto.CreateEnrollmentRequest) (*entity.Enrollment, error) {
	enrollment, err := s.repository.Create(
		ctx,
		&entity.Enrollment{
			NIM:      req.NIM,
			Course:   req.Course,
			Semester: req.Semester,
		},
	)

	if err != nil {
		return nil, err
	}

	return enrollment, nil
}

func (s *enrollmentServiceImpl) Update(ctx context.Context, nim string, course string, req *dto.UpdateEnrollmentRequest) error {
	return s.repository.Update(
		ctx,
		&entity.Enrollment{
			NIM:      nim,
			Course:   course,
			Semester: req.Semester,
			Grade:    req.Grade,
		},
	)
}

func (s *enrollmentServiceImpl) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}

func (s *enrollmentServiceImpl) FindDetailByNim(ctx context.Context, nim string) (*[]entity.EnrollmentDetail, error) {
	return s.repository.FindDetailByNim(ctx, nim)
}
