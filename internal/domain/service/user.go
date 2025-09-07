package service

import (
	"context"
	"fmt"

	"github.com/bobchopperz/bahrululum/internal/domain/models"
	"github.com/bobchopperz/bahrululum/internal/domain/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(ctx context.Context, req *models.CreateUserRequest) (*models.UserResponse, error)
	GetUser(ctx context.Context, id uuid.UUID) (*models.UserResponse, error)
	GetUsers(ctx context.Context, offset, limit int) ([]*models.UserResponse, error)
	UpdateUser(ctx context.Context, id uuid.UUID, updates map[string]interface{}) (*models.UserResponse, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, req *models.CreateUserRequest) (*models.UserResponse, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Nip:      req.Nip,
		Role:     req.Role,
		Password: string(hashPassword),
		IsActive: true,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user.ToResponse(), nil
}

func (s *userService) GetUser(ctx context.Context, id uuid.UUID) (*models.UserResponse, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user.ToResponse(), nil
}

func (s *userService) GetUsers(ctx context.Context, offset, limit int) ([]*models.UserResponse, error) {
	users, err := s.repo.List(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	responses := make([]*models.UserResponse, len(users))
	for i, user := range users {
		responses[i] = user.ToResponse()
	}

	return responses, nil
}

func (s *userService) UpdateUser(ctx context.Context, id uuid.UUID, updates map[string]interface{}) (*models.UserResponse, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if name, ok := updates["name"]; ok {
		user.Name = name.(string)
	}

	if err := s.repo.Update(ctx, user); err != nil {
		return nil, err
	}

	return user.ToResponse(), nil
}

func (s *userService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
