package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VadeMecum struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Description string         `json:"description"`
	Content     string         `json:"content"`
	FileURL     string         `json:"file_url"`
	Category    string         `gorm:"not null" json:"category"`
	Header      string         `json:"cabecalho"`
	TitleID     string         `json:"idtitulo"`
	TitleName   string         `json:"titulo"`
	TitleText   string         `json:"textodotitulo"`
	ChapterID   string         `json:"idcapitulo"`
	ChapterName string         `json:"capitulo"`
	ChapterText string         `json:"textocapitulo"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (v *VadeMecum) BeforeCreate(tx *gorm.DB) error {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return nil
}

type CreateVadeMecumRequest struct {
	Title       string `json:"title" binding:"required,min=3"`
	Description string `json:"description" binding:"required"`
	Content     string `json:"content"`
	FileURL     string `json:"file_url"`
	Category    string `json:"category" binding:"required"`
	Header      string `json:"cabecalho" binding:"required"`
	TitleID     string `json:"idtitulo" binding:"required"`
	TitleName   string `json:"titulo" binding:"required"`
	TitleText   string `json:"textodotitulo" binding:"required"`
	ChapterID   string `json:"idcapitulo" binding:"required"`
	ChapterName string `json:"capitulo" binding:"required"`
	ChapterText string `json:"textocapitulo" binding:"required"`
}

type UpdateVadeMecumRequest struct {
	Title       string `json:"title" binding:"omitempty,min=3"`
	Description string `json:"description"`
	Content     string `json:"content"`
	FileURL     string `json:"file_url"`
	Category    string `json:"category"`
	Header      string `json:"cabecalho"`
	TitleID     string `json:"idtitulo"`
	TitleName   string `json:"titulo"`
	TitleText   string `json:"textodotitulo"`
	ChapterID   string `json:"idcapitulo"`
	ChapterName string `json:"capitulo"`
	ChapterText string `json:"textocapitulo"`
}
