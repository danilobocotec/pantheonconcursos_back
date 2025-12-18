package repository

import (
	"github.com/google/uuid"
	"github.com/thepantheon/api/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type VadeMecumConstituicaoRepository struct {
	db *gorm.DB
}

func NewVadeMecumConstituicaoRepository(db *gorm.DB) *VadeMecumConstituicaoRepository {
	return &VadeMecumConstituicaoRepository{db: db}
}

func (r *VadeMecumConstituicaoRepository) Upsert(items []*model.VadeMecumConstituicao) error {
	if len(items) == 0 {
		return nil
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, item := range items {
			if err := tx.
				Clauses(clause.OnConflict{
					Columns:   []clause.Column{{Name: "registro_id"}},
					UpdateAll: true,
				}).
				Create(item).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *VadeMecumConstituicaoRepository) GetAll() ([]model.VadeMecumConstituicao, error) {
	var items []model.VadeMecumConstituicao
	if err := r.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *VadeMecumConstituicaoRepository) GetByID(id uuid.UUID) (*model.VadeMecumConstituicao, error) {
	var item model.VadeMecumConstituicao
	if err := r.db.First(&item, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *VadeMecumConstituicaoRepository) Create(item *model.VadeMecumConstituicao) error {
	return r.db.Create(item).Error
}

func (r *VadeMecumConstituicaoRepository) Update(item *model.VadeMecumConstituicao) error {
	return r.db.Save(item).Error
}

func (r *VadeMecumConstituicaoRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&model.VadeMecumConstituicao{}, "id = ?", id).Error
}
