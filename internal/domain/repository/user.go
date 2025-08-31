package repository

import (
	"context"

	"github.com/bobchopperz/bahrululum/internal/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, offset, limit int) ([]*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	return nil
}

func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	return nil
}

func (r *userRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}

func (r *userRepository) List(ctx context.Context, offset, limit int) ([]*models.User, error) {
	return nil, nil
}

func (r *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	return nil, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	return nil, nil
}
