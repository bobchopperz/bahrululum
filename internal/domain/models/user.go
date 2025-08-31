package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default;gen_random_uuid()"`
	Name      string         `json:"name" gorm:"not null; size:255" validate:"required,min=2,max=100"`
	Email     string         `json:"email" gorm:"uniqueIndex;not null;size:255" validate:"required,email"`
	Password  string         `json:"-" gorm:"not null;size:255"`
	IsActive  bool           `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,min=1,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) ToResponse() *UserResponse {
	return &UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		IsActive:  u.IsActive,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
