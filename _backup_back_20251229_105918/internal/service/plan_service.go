package service

import (
	"github.com/thepantheon/api/internal/model"
	"github.com/thepantheon/api/internal/repository"
)

type PlanService struct {
	repo *repository.PlanRepository
}

func NewPlanService(repo *repository.PlanRepository) *PlanService {
	return &PlanService{repo: repo}
}

func (s *PlanService) GetPlansWithUsers() ([]model.Plan, error) {
	return s.repo.GetAllWithUsers()
}
