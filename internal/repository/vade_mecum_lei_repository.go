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
