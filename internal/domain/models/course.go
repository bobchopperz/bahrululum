package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Course struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string         `json:"name" gorm:"not null; size:255" validate:"required,min=2,max=100"`
	Description string         `json:"description" gorm:"type:text; not null;"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type CreateCourseRequest struct {
	Name        string `json:"name" validate:"required,min=1,max=100"`
	Description string `json:"description" validate:"required,email"`
}

type CourseResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (u *Course) ToResponse() *CourseResponse {
	return &CourseResponse{
		ID:          u.ID,
		Name:        u.Name,
		Description: u.Description,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}
