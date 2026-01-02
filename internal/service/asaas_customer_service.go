package service

import (
	"github.com/thepantheon/api/internal/model"
	"github.com/thepantheon/api/internal/repository"
)

type AsaasCustomerService struct {
	repo *repository.AsaasCustomerRepository
}

func NewAsaasCustomerService(repo *repository.AsaasCustomerRepository) *AsaasCustomerService {
	return &AsaasCustomerService{repo: repo}
}

func (s *AsaasCustomerService) Create(customer *model.AsaasCustomer) error {
	return s.repo.Create(customer)
}

func (s *AsaasCustomerService) GetByAsaasID(asaasID string) (*model.AsaasCustomer, error) {
	return s.repo.GetByAsaasID(asaasID)
}
