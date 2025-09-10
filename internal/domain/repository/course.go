package repository

import (
	"context"
	"errors"

	"github.com/bobchopperz/bahrululum/internal/domain/models"
	"gorm.io/gorm"
)

type CourseRepository interface {
	Create(ctx context.Context, course *models.Course) error
	GetByID(ctx context.Context, id uint) (*models.Course, error)
	Update(ctx context.Context, course *models.Course) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, offset, limit int) ([]*models.Course, error)
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db}
}

func (r *courseRepository) Create(ctx context.Context, course *models.Course) error {
	if err := r.db.WithContext(ctx).Create(course).Error; err != nil {
		return err
	}
	return nil
}

func (r *courseRepository) Update(ctx context.Context, course *models.Course) error {
	if err := r.db.WithContext(ctx).Save(course).Error; err != nil {
		return err
	}
	return nil
}

func (r *courseRepository) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&models.Course{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (r *courseRepository) List(ctx context.Context, offset, limit int) ([]*models.Course, error) {
	var courses []*models.Course
	err := r.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&courses).Error
	return courses, err
}

func (r *courseRepository) GetByID(ctx context.Context, id uint) (*models.Course, error) {
	var course models.Course
	err := r.db.WithContext(ctx).First(&course, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &course, err
}
