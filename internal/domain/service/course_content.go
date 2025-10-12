package service

import (
	"context"

	"github.com/bobchopperz/bahrululum/internal/domain/models"
	"github.com/bobchopperz/bahrululum/internal/domain/repository"
)

type CourseContentService interface {
	CreateContent(ctx context.Context, req *models.CreateCourseContentRequest) (*models.CourseContentResponse, error)
	GetContent(ctx context.Context, id uint) (*models.CourseContentResponse, error)
	GetContentsByChapter(ctx context.Context, chapterID uint) ([]models.CourseContentResponse, error)
	UpdateContent(ctx context.Context, id uint, req *models.UpdateCourseContentRequest) (*models.CourseContentResponse, error)
	DeleteContent(ctx context.Context, id uint) error
}

type courseContentService struct {
	repo repository.CourseContentRepository
}

func NewCourseContentService(repo repository.CourseContentRepository) CourseContentService {
	return &courseContentService{repo: repo}
}

func (s *courseContentService) CreateContent(ctx context.Context, req *models.CreateCourseContentRequest) (*models.CourseContentResponse, error) {
	content := &models.CourseContent{
		ChapterID:       req.ChapterID,
		Title:           req.Title,
		Description:     req.Description,
		ContentType:     req.ContentType,
		FileURL:         req.FileURL,
		ContentText:     req.ContentText,
		ContentOrder:    req.ContentOrder,
		IsPublished:     req.IsPublished,
		DurationMinutes: req.DurationMinutes,
	}

	if content.ContentOrder == 0 {
		content.ContentOrder = 1
	}

	if err := s.repo.Create(ctx, content); err != nil {
		return nil, err
	}

	return content.ToResponse(), nil
}

func (s *courseContentService) GetContent(ctx context.Context, id uint) (*models.CourseContentResponse, error) {
	content, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return content.ToResponse(), nil
}

func (s *courseContentService) GetContentsByChapter(ctx context.Context, chapterID uint) ([]models.CourseContentResponse, error) {
	contents, err := s.repo.GetByChapterID(ctx, chapterID)
	if err != nil {
		return nil, err
	}

	responses := make([]models.CourseContentResponse, len(contents))
	for i, content := range contents {
		responses[i] = *content.ToResponse()
	}

	return responses, nil
}

func (s *courseContentService) UpdateContent(ctx context.Context, id uint, req *models.UpdateCourseContentRequest) (*models.CourseContentResponse, error) {
	content, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Title != nil {
		content.Title = *req.Title
	}
	if req.Description != nil {
		content.Description = req.Description
	}
	if req.ContentType != nil {
		content.ContentType = *req.ContentType
	}
	if req.FileURL != nil {
		content.FileURL = req.FileURL
	}
	if req.ContentText != nil {
		content.ContentText = req.ContentText
	}
	if req.ContentOrder != nil {
		content.ContentOrder = *req.ContentOrder
	}
	if req.IsPublished != nil {
		content.IsPublished = *req.IsPublished
	}
	if req.DurationMinutes != nil {
		content.DurationMinutes = req.DurationMinutes
	}

	if err := s.repo.Update(ctx, content); err != nil {
		return nil, err
	}

	return content.ToResponse(), nil
}

func (s *courseContentService) DeleteContent(ctx context.Context, id uint) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s *courseContentService) GetContentsByType(ctx context.Context, contentType string) ([]models.CourseContentResponse, error) {
	contents, err := s.repo.GetByContentType(ctx, contentType)
	if err != nil {
		return nil, err
	}

	responses := make([]models.CourseContentResponse, len(contents))
	for i, content := range contents {
		responses[i] = *content.ToResponse()
	}

	return responses, nil
}
