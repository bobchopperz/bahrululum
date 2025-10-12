package service

import (
	"context"

	"github.com/bobchopperz/bahrululum/internal/domain/models"
	"github.com/bobchopperz/bahrululum/internal/domain/repository"
)

type CourseChapterService interface {
	CreateChapter(ctx context.Context, req *models.CreateCourseChapterRequest) (*models.CourseChapterResponse, error)
	GetChapter(ctx context.Context, id uint) (*models.CourseChapterResponse, error)
	GetChaptersByCourse(ctx context.Context, courseID uint) ([]models.CourseChapterResponse, error)
	UpdateChapter(ctx context.Context, id uint, req *models.UpdateCourseChapterRequest) (*models.CourseChapterResponse, error)
	DeleteChapter(ctx context.Context, id uint) error
	GetChapterWithContents(ctx context.Context, id uint) (*models.CourseChapter, error)
}

type courseChapterService struct {
	repo repository.CourseChapterRepository
}

func NewCourseChapterService(repo repository.CourseChapterRepository) CourseChapterService {
	return &courseChapterService{repo: repo}
}

func (s *courseChapterService) CreateChapter(ctx context.Context, req *models.CreateCourseChapterRequest) (*models.CourseChapterResponse, error) {
	chapter := &models.CourseChapter{
		CourseID:     req.CourseID,
		Title:        req.Title,
		Description:  req.Description,
		ChapterOrder: req.ChapterOrder,
		IsPublished:  req.IsPublished,
	}

	if chapter.ChapterOrder == 0 {
		chapter.ChapterOrder = 1
	}

	if err := s.repo.Create(ctx, chapter); err != nil {
		return nil, err
	}

	return chapter.ToResponse(), nil
}

func (s *courseChapterService) GetChapter(ctx context.Context, id uint) (*models.CourseChapterResponse, error) {
	chapter, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return chapter.ToResponse(), nil
}

func (s *courseChapterService) GetChaptersByCourse(ctx context.Context, courseID uint) ([]models.CourseChapterResponse, error) {
	chapters, err := s.repo.GetByCourseID(ctx, courseID)
	if err != nil {
		return nil, err
	}

	responses := make([]models.CourseChapterResponse, len(chapters))
	for i, chapter := range chapters {
		responses[i] = *chapter.ToResponse()
	}

	return responses, nil
}

func (s *courseChapterService) UpdateChapter(ctx context.Context, id uint, req *models.UpdateCourseChapterRequest) (*models.CourseChapterResponse, error) {
	chapter, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Title != nil {
		chapter.Title = *req.Title
	}
	if req.Description != nil {
		chapter.Description = req.Description
	}
	if req.ChapterOrder != nil {
		chapter.ChapterOrder = *req.ChapterOrder
	}
	if req.IsPublished != nil {
		chapter.IsPublished = *req.IsPublished
	}

	if err := s.repo.Update(ctx, chapter); err != nil {
		return nil, err
	}

	return chapter.ToResponse(), nil
}

func (s *courseChapterService) DeleteChapter(ctx context.Context, id uint) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s *courseChapterService) GetChapterWithContents(ctx context.Context, id uint) (*models.CourseChapter, error) {
	chapter, err := s.repo.GetWithContents(ctx, id)
	if err != nil {
		return nil, err
	}

	return chapter, nil
}
