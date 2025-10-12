package models

import (
	"gorm.io/gorm"
	"time"
)

type CourseContent struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	ChapterID       uint           `json:"chapter_id" gorm:"not null"`
	Title           string         `json:"title" gorm:"type:varchar(255);not null"`
	Description     *string        `json:"description" gorm:"type:text"`
	ContentType     string         `json:"content_type" gorm:"type:varchar(50);not null"` // 'video', 'text', 'image', 'pdf', 'link', etc.
	FileURL         *string        `json:"file_url" gorm:"type:varchar(500)"`
	ContentText     *string        `json:"content_text" gorm:"type:text"`
	ContentOrder    int            `json:"content_order" gorm:"not null;default:1"`
	IsPublished     bool           `json:"is_published" gorm:"not null;default:false"`
	DurationMinutes *int           `json:"duration_minutes" gorm:"default:0"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"-" gorm:"index"`

	Chapter CourseChapter `json:"chapter,omitempty" gorm:"foreignKey:ChapterID"`
}

type CreateCourseContentRequest struct {
	ChapterID       uint    `json:"chapter_id" validate:"required"`
	Title           string  `json:"title" validate:"required,min=1,max=255"`
	Description     *string `json:"description,omitempty"`
	ContentType     string  `json:"content_type" validate:"required,oneof=video text image pdf link audio document"`
	FileURL         *string `json:"file_url,omitempty" validate:"omitempty,url"`
	ContentText     *string `json:"content_text,omitempty"`
	ContentOrder    int     `json:"content_order,omitempty"`
	IsPublished     bool    `json:"is_published,omitempty"`
	DurationMinutes *int    `json:"duration_minutes,omitempty"`
}

type UpdateCourseContentRequest struct {
	Title           *string `json:"title,omitempty" validate:"omitempty,min=1,max=255"`
	Description     *string `json:"description,omitempty"`
	ContentType     *string `json:"content_type,omitempty" validate:"omitempty,oneof=video text image pdf link audio document"`
	FileURL         *string `json:"file_url,omitempty" validate:"omitempty,url"`
	ContentText     *string `json:"content_text,omitempty"`
	ContentOrder    *int    `json:"content_order,omitempty"`
	IsPublished     *bool   `json:"is_published,omitempty"`
	DurationMinutes *int    `json:"duration_minutes,omitempty"`
}

type CourseContentResponse struct {
	ID              uint      `json:"id"`
	ChapterID       uint      `json:"chapter_id"`
	Title           string    `json:"title"`
	Description     *string   `json:"description"`
	ContentType     string    `json:"content_type"`
	FileURL         *string   `json:"file_url"`
	ContentText     *string   `json:"content_text"`
	ContentOrder    int       `json:"content_order"`
	IsPublished     bool      `json:"is_published"`
	DurationMinutes *int      `json:"duration_minutes"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (c *CourseContent) ToResponse() *CourseContentResponse {
	return &CourseContentResponse{
		ID:              c.ID,
		ChapterID:       c.ChapterID,
		Title:           c.Title,
		Description:     c.Description,
		ContentType:     c.ContentType,
		FileURL:         c.FileURL,
		ContentText:     c.ContentText,
		ContentOrder:    c.ContentOrder,
		IsPublished:     c.IsPublished,
		DurationMinutes: c.DurationMinutes,
		CreatedAt:       c.CreatedAt,
		UpdatedAt:       c.UpdatedAt,
	}
}
