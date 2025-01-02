package repository

import (
	"github.com/thepantheon/api/internal/model"
	"gorm.io/gorm"
)

type MediaAssetRepository struct {
	db *gorm.DB
}

func NewMediaAssetRepository(db *gorm.DB) *MediaAssetRepository {
	return &MediaAssetRepository{db: db}
}

func (r *MediaAssetRepository) Create(item *model.MediaAsset) error {
	return r.db.Create(item).Error
}

func (r *MediaAssetRepository) GetByID(id string) (*model.MediaAsset, error) {
	var item model.MediaAsset
	if err := r.db.First(&item, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MediaAssetRepository) Delete(id string) error {
	return r.db.Delete(&model.MediaAsset{}, "id = ?", id).Error
}
