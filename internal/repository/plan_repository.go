package repository

import (
	"github.com/thepantheon/api/internal/model"
	"gorm.io/gorm"
)

type PlanRepository struct {
	db *gorm.DB
}

func NewPlanRepository(db *gorm.DB) *PlanRepository {
	return &PlanRepository{db: db}
}

func (r *PlanRepository) GetAllWithUsers() ([]model.Plan, error) {
	var plans []model.Plan
	if err := r.db.Preload("Users").Find(&plans).Error; err != nil {
		return nil, err
	}
	return plans, nil
}
