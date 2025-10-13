package service

import (
	"context"

	"github.com/bobchopperz/bahrululum/internal/domain/models"
	"github.com/bobchopperz/bahrululum/internal/domain/repository"
)

type EnrollmentService interface {
	Create(ctx context.Context, userID uint, req *models.CreateEnrollmentRequest) (*models.EnrollmentResponse, error)
	GetByCouseID(ctx context.Context, id uint) (*models.EnrollmentResponse, error)
	CheckEnrollment(ctx context.Context, userID, courseID uint) (bool, error)
	GetUserEnrollments(ctx context.Context, userID uint) ([]uint, error)
}

type enrollmentService struct {
	repo repository.EnrollmentRepository
}

func NewEnrollmentService(repo repository.EnrollmentRepository) EnrollmentService {
	return &enrollmentService{repo: repo}
}

func (s *enrollmentService) Create(ctx context.Context, userID uint, req *models.CreateEnrollmentRequest) (*models.EnrollmentResponse, error) {
	existing, err := s.repo.GetByUserAndCourse(ctx, userID, req.CourseID)
	if err != nil {
		return nil, err
	}

	if existing != nil {
		return existing.ToResponse(), nil
	}

	enrollment := &models.Enrollment{
		CourseID: req.CourseID,
		UserID:   userID,
	}

	if err := s.repo.Create(ctx, enrollment); err != nil {
		return nil, err
	}

	return enrollment.ToResponse(), nil
}

func (s *enrollmentService) CheckEnrollment(ctx context.Context, userID, courseID uint) (bool, error) {
	enrollment, err := s.repo.GetByUserAndCourse(ctx, userID, courseID)
	if err != nil {
		return false, err
	}
	return enrollment != nil, nil
}

func (s *enrollmentService) GetUserEnrollments(ctx context.Context, userID uint) ([]uint, error) {
	enrollments, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	courseIDs := make([]uint, len(enrollments))
	for i, enrollment := range enrollments {
		courseIDs[i] = enrollment.CourseID
	}

	return courseIDs, nil
}

func (s *enrollmentService) GetByCouseID(ctx context.Context, id uint) (*models.EnrollmentResponse, error) {
	entity, err := s.repo.GetByCourseID(ctx, id)
	if err != nil {
		return nil, err
	}

	return entity.ToResponse(), nil
}
