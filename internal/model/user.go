package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID         uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Email      string         `gorm:"uniqueIndex;not null" json:"email"`
	FullName   string         `gorm:"not null" json:"full_name"`
	Password   string         `json:"-"`
	Avatar     string         `json:"avatar,omitempty"`
	Provider   string         `json:"provider,omitempty"` // local, google, facebook
	ProviderID string         `json:"provider_id,omitempty"`
	PlanID     *uuid.UUID     `gorm:"type:uuid" json:"plan_id,omitempty"`
	Plan       *Plan          `gorm:"foreignKey:PlanID" json:"plan,omitempty"`
	Active     bool           `gorm:"default:true" json:"active"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

// BeforeCreate generates a UUID for the user if not already set
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	FullName string `json:"full_name" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=8"`
	Confirm  string `json:"confirm" binding:"required,eqfield=Password"`
}

type UpdateUserRequest struct {
	Email    string `json:"email" binding:"email"`
	FullName string `json:"full_name"`
	Active   bool   `json:"active"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token     string `json:"token"`
	ExpiresIn int64  `json:"expires_in"`
	User      User   `json:"user"`
}

type RefreshTokenRequest struct {
	Token string `json:"token" binding:"required"`
}

type SocialAuthRequest struct {
	Provider    string `json:"provider" binding:"required,oneof=google facebook"`
	AccessToken string `json:"access_token" binding:"required"`
}

type SocialUserInfo struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Picture  string `json:"picture"`
	Provider string `json:"provider"`
}
