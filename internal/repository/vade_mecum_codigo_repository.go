package repository

import (
	"github.com/google/uuid"
	"github.com/thepantheon/api/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type VadeMecumCodigoRepository struct {
	db *gorm.DB
}

func NewVadeMecumCodigoRepository(db *gorm.DB) *VadeMecumCodigoRepository {
	return &VadeMecumCodigoRepository{db: db}
}

func (r *VadeMecumCodigoRepository) Create(item *model.VadeMecumCodigo) error {
	return r.db.Create(item).Error
}

func (r *VadeMecumCodigoRepository) CreateBatch(items []*model.VadeMecumCodigo) error {
	if len(items) == 0 {
		return nil
	}
	return r.db.Create(&items).Error
}

func (r *VadeMecumCodigoRepository) UpsertByCodigo(items []*model.VadeMecumCodigo) error {
	if len(items) == 0 {
		return nil
	}
	return r.db.
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "idcodigo"}},
			UpdateAll: true,
		}).
		Create(&items).Error
}

func (r *VadeMecumCodigoRepository) GetAll() ([]model.VadeMecumCodigo, error) {
	var items []model.VadeMecumCodigo
	if err := r.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *VadeMecumCodigoRepository) GetByID(id uuid.UUID) (*model.VadeMecumCodigo, error) {
	var item model.VadeMecumCodigo
	if err := r.db.First(&item, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *VadeMecumCodigoRepository) Update(item *model.VadeMecumCodigo) error {
	return r.db.Save(item).Error
}

func (r *VadeMecumCodigoRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&model.VadeMecumCodigo{}, "id = ?", id).Error
}
