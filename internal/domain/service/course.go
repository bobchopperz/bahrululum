package service

import (
	"context"

	"github.com/bobchopperz/bahrululum/internal/domain/models"
	"github.com/bobchopperz/bahrululum/internal/domain/repository"
)

type CourseService interface {
	CreateCourse(ctx context.Context, req *models.CreateCourseRequest) (*models.CourseResponse, error)
	GetCourse(ctx context.Context, id uint) (*models.CourseResponse, error)
	GetCourses(ctx context.Context, offset, limit int) ([]*models.CourseResponse, error)
	UpdateCourse(ctx context.Context, id uint, updates map[string]interface{}) (*models.CourseResponse, error)
	DeleteCourse(ctx context.Context, id uint) error
}

type courseService struct {
	repo repository.CourseRepository
}

func NewCourseService(repo repository.CourseRepository) CourseService {
	return &courseService{repo: repo}
}

func (s *courseService) CreateCourse(ctx context.Context, req *models.CreateCourseRequest) (*models.CourseResponse, error) {
	course := &models.Course{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := s.repo.Create(ctx, course); err != nil {
		return nil, err
	}

	return course.ToResponse(), nil
}

func (s *courseService) GetCourse(ctx context.Context, id uint) (*models.CourseResponse, error) {
	course, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return course.ToResponse(), nil
}

func (s *courseService) GetCourses(ctx context.Context, offset, limit int) ([]*models.CourseResponse, error) {
	courses, err := s.repo.List(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	responses := make([]*models.CourseResponse, len(courses))
	for i, course := range courses {
		responses[i] = course.ToResponse()
	}

	return responses, nil
}

func (s *courseService) UpdateCourse(ctx context.Context, id uint, updates map[string]interface{}) (*models.CourseResponse, error) {
	course, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if name, ok := updates["name"]; ok {
		course.Name = name.(string)
	}

	if err := s.repo.Update(ctx, course); err != nil {
		return nil, err
	}

	return course.ToResponse(), nil
}

func (s *courseService) DeleteCourse(ctx context.Context, id uint) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
