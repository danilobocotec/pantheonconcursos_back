package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Plan struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string    `gorm:"uniqueIndex;not null" json:"name"`
	Description string    `json:"description"`
	Phase       string    `json:"phase"`
	Duration    string    `json:"duration"`
	Active      bool      `gorm:"default:true" json:"active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Users       []User    `gorm:"foreignKey:PlanID" json:"users,omitempty"`
}

// BeforeCreate ensures UUID is set for plan.
func (p *Plan) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}
