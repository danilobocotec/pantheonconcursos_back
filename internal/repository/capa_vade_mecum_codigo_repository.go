package repository

import (
	"github.com/thepantheon/api/internal/model"
	"gorm.io/gorm"
)

type CapaVadeMecumCodigoRepository struct {
	db *gorm.DB
}

func NewCapaVadeMecumCodigoRepository(db *gorm.DB) *CapaVadeMecumCodigoRepository {
	return &CapaVadeMecumCodigoRepository{db: db}
}

func (r *CapaVadeMecumCodigoRepository) GetAll() ([]model.CapaVadeMecumCodigo, error) {
	var items []model.CapaVadeMecumCodigo
	if err := r.db.Order("nomecodigo ASC").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *CapaVadeMecumCodigoRepository) GetByNomeCodigo(nomecodigo string) (*model.CapaVadeMecumCodigo, error) {
	var item model.CapaVadeMecumCodigo
	if err := r.db.First(&item, "nomecodigo = ?", nomecodigo).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *CapaVadeMecumCodigoRepository) Create(item *model.CapaVadeMecumCodigo) error {
	return r.db.Create(item).Error
}

func (r *CapaVadeMecumCodigoRepository) Update(item *model.CapaVadeMecumCodigo) error {
	return r.db.Save(item).Error
}
