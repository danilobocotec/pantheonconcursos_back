package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AsaasPayment struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	AsaasID      string    `gorm:"uniqueIndex;not null" json:"asaas_id"`
	CustomerID   string    `gorm:"not null" json:"customer_id"`
	Value        float64   `gorm:"not null" json:"value"`
	DueDate      string    `gorm:"not null" json:"due_date"`
	BillingType  string    `gorm:"not null" json:"billing_type"`
	RequestJSON  string    `gorm:"type:text" json:"request_json"`
	ResponseJSON string    `gorm:"type:text" json:"response_json"`
	ConfirmationJSON string `gorm:"type:text" json:"confirmation_json"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// BeforeCreate ensures UUID is set for Asaas payment.
func (p *AsaasPayment) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}
