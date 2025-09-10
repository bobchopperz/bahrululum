package models

import (
	"time"

	"github.com/bobchopperz/bahrululum/internal/constants"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string         `json:"name" gorm:"not null; size:255" validate:"required,min=2,max=100"`
	Email     string         `json:"email" gorm:"uniqueIndex;not null;size:255" validate:"required,email"`
	Nip       string         `json:"nip" gorm:"uniqueIndex;not null;size:12" validate:"required,min=12,max=12"`
	Password  string         `json:"-" gorm:"not null;size:255"`
	IsActive  bool           `json:"is_active" gorm:"default:true"`
	Role      string         `json:"role" gorm:"not null;size:50;default:'user'" validate:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,min=1,max=100"`
	Nip      string `json:"nip" validate:"required,min=12"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

type LoginRequest struct {
	Nip      string `json:"nip" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	ID        uint `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	IsActive  bool      `json:"is_active"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) ToResponse() *UserResponse {
	return &UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		IsActive:  u.IsActive,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (u *User) GetRole() constants.Role {
	role, _ := constants.ParseRole(u.Role)
	return role
}

func (u *User) SetRole(role constants.Role) {
	u.Role = role.String()
}

func (u *User) HasRole(role constants.Role) bool {
	return u.Role == role.String()
}

func (u *User) IsAdmin() bool {
	return u.HasRole(constants.RoleAdmin)
}

func (u *User) IsMentor() bool {
	return u.HasRole(constants.RoleMentor)
}

func (u *User) IsUser() bool {
	return u.HasRole(constants.RoleUser)
}
