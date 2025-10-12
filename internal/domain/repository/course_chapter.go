package repository

import (
	"context"
	"errors"

	"github.com/bobchopperz/bahrululum/internal/domain/models"
	"gorm.io/gorm"
)

type CourseChapterRepository interface {
	Create(ctx context.Context, chapter *models.CourseChapter) error
	GetByID(ctx context.Context, id uint) (*models.CourseChapter, error)
	GetByCourseID(ctx context.Context, courseID uint) ([]models.CourseChapter, error)
	Update(ctx context.Context, chapter *models.CourseChapter) error
	Delete(ctx context.Context, id uint) error
	GetWithContents(ctx context.Context, id uint) (*models.CourseChapter, error)
}

type courseChapterRepository struct {
	db *gorm.DB
}

func NewCourseChapterRepository(db *gorm.DB) CourseChapterRepository {
	return &courseChapterRepository{db}
}

func (r *courseChapterRepository) Create(ctx context.Context, chapter *models.CourseChapter) error {
	if err := r.db.WithContext(ctx).Create(chapter).Error; err != nil {
		return err
	}
	return nil
}

func (r *courseChapterRepository) GetByID(ctx context.Context, id uint) (*models.CourseChapter, error) {
	var chapter models.CourseChapter
	err := r.db.WithContext(ctx).First(&chapter, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &chapter, err
}

func (r *courseChapterRepository) GetByCourseID(ctx context.Context, courseID uint) ([]models.CourseChapter, error) {
	var chapters []models.CourseChapter
	err := r.db.WithContext(ctx).Where("course_id = ?", courseID).Order("chapter_order ASC").Find(&chapters).Error
	return chapters, err
}

func (r *courseChapterRepository) Update(ctx context.Context, chapter *models.CourseChapter) error {
	if err := r.db.WithContext(ctx).Save(chapter).Error; err != nil {
		return err
	}
	return nil
}

func (r *courseChapterRepository) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&models.CourseChapter{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (r *courseChapterRepository) GetWithContents(ctx context.Context, id uint) (*models.CourseChapter, error) {
	var chapter models.CourseChapter
	err := r.db.WithContext(ctx).Preload("Contents").First(&chapter, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &chapter, err
}
