package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// MediaAsset stores binary audio/image data and metadata.
type MediaAsset struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Kind        string         `gorm:"type:varchar(20);not null;index" json:"tipo"`
	Filename    string         `gorm:"type:text" json:"nome_arquivo"`
	ContentType string         `gorm:"type:text" json:"content_type"`
	Size        int64          `json:"tamanho"`
	Data        []byte         `gorm:"type:bytea" json:"-"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (MediaAsset) TableName() string {
	return "media_assets"
}

func (m *MediaAsset) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}
