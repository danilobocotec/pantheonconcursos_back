package service

import (
	"github.com/thepantheon/api/internal/model"
	"github.com/thepantheon/api/internal/repository"
)

type AsaasPaymentService struct {
	repo *repository.AsaasPaymentRepository
}

func NewAsaasPaymentService(repo *repository.AsaasPaymentRepository) *AsaasPaymentService {
	return &AsaasPaymentService{repo: repo}
}

func (s *AsaasPaymentService) Create(payment *model.AsaasPayment) error {
	return s.repo.Create(payment)
}

func (s *AsaasPaymentService) GetByAsaasID(asaasID string) (*model.AsaasPayment, error) {
	return s.repo.GetByAsaasID(asaasID)
}

func (s *AsaasPaymentService) UpdateConfirmation(asaasID, confirmationJSON string) error {
	return s.repo.UpdateConfirmation(asaasID, confirmationJSON)
}
