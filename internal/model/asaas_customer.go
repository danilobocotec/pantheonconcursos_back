package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AsaasCustomer struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	AsaasID      string    `gorm:"uniqueIndex;not null" json:"asaas_id"`
	Name         string    `gorm:"not null" json:"name"`
	CPFOrCNPJ    string    `gorm:"not null" json:"cpf_cnpj"`
	Email        string    `gorm:"not null" json:"email"`
	Phone        string    `gorm:"not null" json:"phone"`
	ResponseJSON string    `gorm:"type:text" json:"response_json"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// BeforeCreate ensures UUID is set for Asaas customer.
func (c *AsaasCustomer) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}
