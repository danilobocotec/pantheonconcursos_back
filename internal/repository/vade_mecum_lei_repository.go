package repository

import (
	"github.com/thepantheon/api/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type VadeMecumLeiRepository struct {
	db *gorm.DB
}

func NewVadeMecumLeiRepository(db *gorm.DB) *VadeMecumLeiRepository {
	return &VadeMecumLeiRepository{db: db}
}

func (r *VadeMecumLeiRepository) Upsert(items []*model.VadeMecumLei) error {
	if len(items) == 0 {
		return nil
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, item := range items {
			if err := tx.
				Clauses(clause.OnConflict{
					Columns:   []clause.Column{{Name: "id"}},
					UpdateAll: true,
				}).
				Create(item).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *VadeMecumLeiRepository) Create(item *model.VadeMecumLei) error {
	return r.db.Create(item).Error
}

func (r *VadeMecumLeiRepository) GetAll() ([]model.VadeMecumLei, error) {
	var items []model.VadeMecumLei
	if err := r.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *VadeMecumLeiRepository) GetByID(id string) (*model.VadeMecumLei, error) {
	var item model.VadeMecumLei
	if err := r.db.First(&item, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *VadeMecumLeiRepository) Update(item *model.VadeMecumLei) error {
	return r.db.Save(item).Error
}

func (r *VadeMecumLeiRepository) Delete(id string) error {
	return r.db.Delete(&model.VadeMecumLei{}, "id = ?", id).Error
}
