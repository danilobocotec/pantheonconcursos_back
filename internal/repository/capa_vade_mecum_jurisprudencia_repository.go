package repository

import (
	"github.com/thepantheon/api/internal/model"
	"gorm.io/gorm"
)

type CapaVadeMecumJurisprudenciaRepository struct {
	db *gorm.DB
}

func NewCapaVadeMecumJurisprudenciaRepository(db *gorm.DB) *CapaVadeMecumJurisprudenciaRepository {
	return &CapaVadeMecumJurisprudenciaRepository{db: db}
}

func (r *CapaVadeMecumJurisprudenciaRepository) GetAll() ([]model.CapaVadeMecumJurisprudencia, error) {
	var items []model.CapaVadeMecumJurisprudencia
	if err := r.db.Order("nomecodigo ASC").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *CapaVadeMecumJurisprudenciaRepository) GetByNomeCodigo(nomecodigo string) (*model.CapaVadeMecumJurisprudencia, error) {
	var item model.CapaVadeMecumJurisprudencia
	if err := r.db.First(&item, "nomecodigo = ?", nomecodigo).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *CapaVadeMecumJurisprudenciaRepository) Create(item *model.CapaVadeMecumJurisprudencia) error {
	return r.db.Create(item).Error
}

func (r *CapaVadeMecumJurisprudenciaRepository) Update(item *model.CapaVadeMecumJurisprudencia) error {
	return r.db.Save(item).Error
}
