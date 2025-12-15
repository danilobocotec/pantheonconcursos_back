package repository

import (
	"github.com/thepantheon/api/internal/model"
	"gorm.io/gorm"
)

type CapaVadeMecumOABRepository struct {
	db *gorm.DB
}

func NewCapaVadeMecumOABRepository(db *gorm.DB) *CapaVadeMecumOABRepository {
	return &CapaVadeMecumOABRepository{db: db}
}

func (r *CapaVadeMecumOABRepository) GetAll() ([]model.CapaVadeMecumOAB, error) {
	var items []model.CapaVadeMecumOAB
	if err := r.db.Order("nomecodigo ASC").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *CapaVadeMecumOABRepository) GetByNomeCodigo(nomecodigo string) (*model.CapaVadeMecumOAB, error) {
	var item model.CapaVadeMecumOAB
	if err := r.db.First(&item, "nomecodigo = ?", nomecodigo).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *CapaVadeMecumOABRepository) Create(item *model.CapaVadeMecumOAB) error {
	return r.db.Create(item).Error
}

func (r *CapaVadeMecumOABRepository) Update(item *model.CapaVadeMecumOAB) error {
	return r.db.Save(item).Error
}
