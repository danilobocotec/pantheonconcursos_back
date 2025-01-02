package repository

import (
	"github.com/thepantheon/api/internal/model"
	"gorm.io/gorm"
)

type AsaasCustomerRepository struct {
	db *gorm.DB
}

func NewAsaasCustomerRepository(db *gorm.DB) *AsaasCustomerRepository {
	return &AsaasCustomerRepository{db: db}
}

func (r *AsaasCustomerRepository) Create(customer *model.AsaasCustomer) error {
	return r.db.Create(customer).Error
}

func (r *AsaasCustomerRepository) GetByAsaasID(asaasID string) (*model.AsaasCustomer, error) {
	var customer model.AsaasCustomer
	if err := r.db.Where("asaas_id = ?", asaasID).First(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}
