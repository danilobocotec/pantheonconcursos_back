package repository

import (
	"github.com/google/uuid"
	"github.com/thepantheon/api/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type VadeMecumEstatutoRepository struct {
	db *gorm.DB
}

func NewVadeMecumEstatutoRepository(db *gorm.DB) *VadeMecumEstatutoRepository {
	return &VadeMecumEstatutoRepository{db: db}
}

func (r *VadeMecumEstatutoRepository) UpsertByCodigo(items []*model.VadeMecumEstatuto) error {
	if len(items) == 0 {
		return nil
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, item := range items {
			if err := tx.
				Clauses(clause.OnConflict{
					Columns:   []clause.Column{{Name: "idcodigo"}},
					UpdateAll: true,
				}).
				Create(item).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *VadeMecumEstatutoRepository) DeleteAll() error {
	return r.db.Where("1=1").Delete(&model.VadeMecumEstatuto{}).Error
}

func (r *VadeMecumEstatutoRepository) GetByID(id uuid.UUID) (*model.VadeMecumEstatuto, error) {
	var item model.VadeMecumEstatuto
	if err := r.db.First(&item, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *VadeMecumEstatutoRepository) GetAll() ([]model.VadeMecumEstatuto, error) {
	var items []model.VadeMecumEstatuto
	if err := r.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *VadeMecumEstatutoRepository) Create(item *model.VadeMecumEstatuto) error {
	return r.db.Create(item).Error
}

func (r *VadeMecumEstatutoRepository) Update(item *model.VadeMecumEstatuto) error {
	return r.db.Save(item).Error
}

func (r *VadeMecumEstatutoRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&model.VadeMecumEstatuto{}, "id = ?", id).Error
}
