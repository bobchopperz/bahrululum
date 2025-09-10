package models

import (
	"gorm.io/gorm"
)

type Enrollment struct {
	gorm.Model
	UserID   uint `json:"user_id"`
	CourseID uint `json:"course_id"` 
	User     User
	Course   Course
}

type CreateEnrollmentRequest struct {
	UserID   string `json:"user_id" validate:"required"`
	CourseID string `json:"course_id" validate:"required"`
}

type EnrollmentResponse struct {
	UserID   string `json:"user_id"`
	CourseID string `json:"course_id"`
}
