package repository

import (
	"github.com/thepantheon/api/internal/model"
	"gorm.io/gorm"
)

type AsaasPaymentRepository struct {
	db *gorm.DB
}

func NewAsaasPaymentRepository(db *gorm.DB) *AsaasPaymentRepository {
	return &AsaasPaymentRepository{db: db}
}

func (r *AsaasPaymentRepository) Create(payment *model.AsaasPayment) error {
	return r.db.Create(payment).Error
}

func (r *AsaasPaymentRepository) GetByAsaasID(asaasID string) (*model.AsaasPayment, error) {
	var payment model.AsaasPayment
	if err := r.db.Where("asaas_id = ?", asaasID).First(&payment).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *AsaasPaymentRepository) UpdateConfirmation(asaasID, confirmationJSON string) error {
	return r.db.Model(&model.AsaasPayment{}).
		Where("asaas_id = ?", asaasID).
		Update("confirmation_json", confirmationJSON).Error
}
