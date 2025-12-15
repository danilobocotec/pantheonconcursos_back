package repository

import (
	"github.com/google/uuid"
	"github.com/thepantheon/api/internal/model"
	"gorm.io/gorm"
)

type VadeMecumRepository struct {
	db *gorm.DB
}

func NewVadeMecumRepository(db *gorm.DB) *VadeMecumRepository {
	return &VadeMecumRepository{db: db}
}

func (r *VadeMecumRepository) Create(vm *model.VadeMecum) error {
	return r.db.Create(vm).Error
}

func (r *VadeMecumRepository) GetAll() ([]model.VadeMecum, error) {
	var items []model.VadeMecum
	if err := r.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *VadeMecumRepository) GetByCategory(category string) ([]model.VadeMecum, error) {
	var items []model.VadeMecum
	if err := r.db.Where("category = ?", category).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *VadeMecumRepository) GetByID(id uuid.UUID) (*model.VadeMecum, error) {
	var vm model.VadeMecum
	if err := r.db.First(&vm, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &vm, nil
}

func (r *VadeMecumRepository) Update(vm *model.VadeMecum) error {
	return r.db.Save(vm).Error
}

func (r *VadeMecumRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&model.VadeMecum{}, "id = ?", id).Error
}
