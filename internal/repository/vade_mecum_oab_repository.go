package repository

import (
	"github.com/thepantheon/api/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type VadeMecumOABRepository struct {
	db *gorm.DB
}

func NewVadeMecumOABRepository(db *gorm.DB) *VadeMecumOABRepository {
	return &VadeMecumOABRepository{db: db}
}

func (r *VadeMecumOABRepository) Upsert(items []*model.VadeMecumOAB) error {
	if len(items) == 0 {
		return nil
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			UpdateAll: true,
		}).Create(items).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *VadeMecumOABRepository) GetAll() ([]model.VadeMecumOAB, error) {
	var results []model.VadeMecumOAB
	if err := r.db.
		Order("nomecodigo ASC").
		Order("num_artigo ASC").
		Order("\"Artigos\" ASC").
		Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (r *VadeMecumOABRepository) DeleteAll() error {
	return r.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.VadeMecumOAB{}).Error
}

func (r *VadeMecumOABRepository) Create(item *model.VadeMecumOAB) error {
	return r.db.Create(item).Error
}

func (r *VadeMecumOABRepository) GetByID(id string) (*model.VadeMecumOAB, error) {
	var item model.VadeMecumOAB
	if err := r.db.First(&item, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *VadeMecumOABRepository) Update(item *model.VadeMecumOAB) error {
	return r.db.Save(item).Error
}

func (r *VadeMecumOABRepository) Delete(id string) error {
	return r.db.Delete(&model.VadeMecumOAB{}, "id = ?", id).Error
}
