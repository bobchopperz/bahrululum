package repository

import (
	"context"
	"errors"

	"github.com/bobchopperz/bahrululum/internal/domain/models"
	"gorm.io/gorm"
)

type CourseContentRepository interface {
	Create(ctx context.Context, content *models.CourseContent) error
	GetByID(ctx context.Context, id uint) (*models.CourseContent, error)
	GetByChapterID(ctx context.Context, chapterID uint) ([]models.CourseContent, error)
	Update(ctx context.Context, content *models.CourseContent) error
	Delete(ctx context.Context, id uint) error
	GetByContentType(ctx context.Context, contentType string) ([]models.CourseContent, error)
}

type courseContentRepository struct {
	db *gorm.DB
}

func NewCourseContentRepository(db *gorm.DB) CourseContentRepository {
	return &courseContentRepository{db}
}

func (r *courseContentRepository) Create(ctx context.Context, content *models.CourseContent) error {
	if err := r.db.WithContext(ctx).Create(content).Error; err != nil {
		return err
	}
	return nil
}

func (r *courseContentRepository) GetByID(ctx context.Context, id uint) (*models.CourseContent, error) {
	var content models.CourseContent
	err := r.db.WithContext(ctx).First(&content, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &content, err
}

func (r *courseContentRepository) GetByChapterID(ctx context.Context, chapterID uint) ([]models.CourseContent, error) {
	var contents []models.CourseContent
	err := r.db.WithContext(ctx).Where("chapter_id = ?", chapterID).Order("content_order ASC").Find(&contents).Error
	return contents, err
}

func (r *courseContentRepository) Update(ctx context.Context, content *models.CourseContent) error {
	if err := r.db.WithContext(ctx).Save(content).Error; err != nil {
		return err
	}
	return nil
}

func (r *courseContentRepository) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&models.CourseContent{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (r *courseContentRepository) GetByContentType(ctx context.Context, contentType string) ([]models.CourseContent, error) {
	var contents []models.CourseContent
	err := r.db.WithContext(ctx).Where("content_type = ?", contentType).Find(&contents).Error
	return contents, err
}
