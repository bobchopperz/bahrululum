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
	UserID   uint `json:"user_id" validate:"required"`
	CourseID uint `json:"course_id" validate:"required"`
}

type EnrollmentResponse struct {
	UserID   uint `json:"user_id"`
	CourseID uint `json:"course_id"`
}

func (u *Enrollment) ToResponse() *EnrollmentResponse {
	return &EnrollmentResponse{
		UserID:   u.UserID,
		CourseID: u.CourseID,
	}
}
