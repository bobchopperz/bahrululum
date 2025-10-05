package models

import (
	"time"

	"gorm.io/gorm"
)

type RefreshToken struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint           `json:"user_id" gorm:"not null"`
	Token     string         `json:"token" gorm:"uniqueIndex;not null;size:500"`
	ExpiresAt time.Time      `json:"expires_at"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	User User `json:"user" gorm:"constraint:OnUpdate:Cascade,OnDelete:CASCADE;"`
}

type TokenResponse struct {
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
	ExpiresAt    int64         `json:"expires_at"`
	TokenType    string        `json:"token_type"`
	User         *UserResponse `json:"user,omitempty"`
}
