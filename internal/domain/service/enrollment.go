package service

import (
	"context"

	"github.com/bobchopperz/bahrululum/internal/domain/models"
	"github.com/bobchopperz/bahrululum/internal/domain/repository"
)

type EnrollmentService interface {
	Create(ctx context.Context, userID uint, req *models.CreateEnrollmentRequest) (*models.EnrollmentResponse, error)
	GetByCouseID(ctx context.Context, id uint) (*models.EnrollmentResponse, error)
}

type enrollmentService struct {
	repo repository.EnrollmentRepository
}

func NewEnrollmentService(repo repository.EnrollmentRepository) EnrollmentService {
	return &enrollmentService{repo: repo}
}

func (s *enrollmentService) Create(ctx context.Context, userID uint, req *models.CreateEnrollmentRequest) (*models.EnrollmentResponse, error) {
	enrollment := &models.Enrollment{
		CourseID: req.CourseID,
		UserID:   userID,
	}

	if err := s.repo.Create(ctx, enrollment); err != nil {
		return nil, err
	}

	return enrollment.ToResponse(), nil
}

func (s *enrollmentService) GetByCouseID(ctx context.Context, id uint) (*models.EnrollmentResponse, error) {
	entity, err := s.repo.GetByCourseID(ctx, id)
	if err != nil {
		return nil, err
	}

	return entity.ToResponse(), nil
}
