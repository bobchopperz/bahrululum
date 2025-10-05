package repository

import (
	"context"
	"errors"

	"github.com/bobchopperz/bahrululum/internal/domain/models"
	"gorm.io/gorm"
)

type EnrollmentRepository interface {
	Create(ctx context.Context, course *models.Enrollment) error
	GetByID(ctx context.Context, id uint) (*models.Enrollment, error)
	GetByCourseID(ctx context.Context, id uint) (*models.Enrollment, error)
	Update(ctx context.Context, course *models.Enrollment) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, offset, limit int) ([]*models.Enrollment, error)
}

type enrollmentRepository struct {
	db *gorm.DB
}

func NewEnrollmentRepository(db *gorm.DB) EnrollmentRepository {
	return &enrollmentRepository{db}
}

func (r *enrollmentRepository) Create(ctx context.Context, course *models.Enrollment) error {
	if err := r.db.WithContext(ctx).Create(course).Error; err != nil {
		return err
	}
	return nil
}

func (r *enrollmentRepository) Update(ctx context.Context, course *models.Enrollment) error {
	if err := r.db.WithContext(ctx).Save(course).Error; err != nil {
		return err
	}
	return nil
}

func (r *enrollmentRepository) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&models.Enrollment{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func (r *enrollmentRepository) List(ctx context.Context, offset, limit int) ([]*models.Enrollment, error) {
	var enrollments []*models.Enrollment
	err := r.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&enrollments).Error
	return enrollments, err
}

func (r *enrollmentRepository) GetByID(ctx context.Context, id uint) (*models.Enrollment, error) {
	var enrollment models.Enrollment
	err := r.db.WithContext(ctx).First(&enrollment, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &enrollment, err
}

func (r *enrollmentRepository) GetByCourseID(ctx context.Context, id uint) (*models.Enrollment, error) {
	var enrollment models.Enrollment
	err := r.db.WithContext(ctx).First(&enrollment, "course_id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &enrollment, err
}
