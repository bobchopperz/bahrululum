package models

import (
	"gorm.io/gorm"
	"time"
)

type CourseChapter struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CourseID     uint           `json:"course_id" gorm:"not null"`
	Title        string         `json:"title" gorm:"type:varchar(255);not null"`
	Description  *string        `json:"description" gorm:"type:text"`
	ChapterOrder int            `json:"chapter_order" gorm:"not null;default:1"`
	IsPublished  bool           `json:"is_published" gorm:"not null;default:false"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`

	Course   Course          `json:"course,omitempty" gorm:"foreignKey:CourseID"`
	Contents []CourseContent `json:"contents,omitempty" gorm:"foreignKey:ChapterID"`
}

type CreateCourseChapterRequest struct {
	CourseID     uint    `json:"course_id" validate:"required"`
	Title        string  `json:"title" validate:"required,min=1,max=255"`
	Description  *string `json:"description,omitempty"`
	ChapterOrder int     `json:"chapter_order,omitempty"`
	IsPublished  bool    `json:"is_published,omitempty"`
}

type UpdateCourseChapterRequest struct {
	Title        *string `json:"title,omitempty" validate:"omitempty,min=1,max=255"`
	Description  *string `json:"description,omitempty"`
	ChapterOrder *int    `json:"chapter_order,omitempty"`
	IsPublished  *bool   `json:"is_published,omitempty"`
}

type CourseChapterResponse struct {
	ID           uint      `json:"id"`
	CourseID     uint      `json:"course_id"`
	Title        string    `json:"title"`
	Description  *string   `json:"description"`
	ChapterOrder int       `json:"chapter_order"`
	IsPublished  bool      `json:"is_published"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (c *CourseChapter) ToResponse() *CourseChapterResponse {
	return &CourseChapterResponse{
		ID:           c.ID,
		CourseID:     c.CourseID,
		Title:        c.Title,
		Description:  c.Description,
		ChapterOrder: c.ChapterOrder,
		IsPublished:  c.IsPublished,
		CreatedAt:    c.CreatedAt,
		UpdatedAt:    c.UpdatedAt,
	}
}
