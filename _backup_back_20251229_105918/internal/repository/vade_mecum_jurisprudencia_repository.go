package repository

import (
	"github.com/thepantheon/api/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type VadeMecumJurisprudenciaRepository struct {
	db *gorm.DB
}

func NewVadeMecumJurisprudenciaRepository(db *gorm.DB) *VadeMecumJurisprudenciaRepository {
	return &VadeMecumJurisprudenciaRepository{db: db}
}

func (r *VadeMecumJurisprudenciaRepository) Upsert(items []*model.VadeMecumJurisprudencia) error {
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

func (r *VadeMecumJurisprudenciaRepository) GetAll() ([]model.VadeMecumJurisprudencia, error) {
	var items []model.VadeMecumJurisprudencia
	if err := r.db.
		Order("nomecodigo ASC").
		Order("\"Normativo\" ASC").
		Order("num_artigo ASC").
		Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *VadeMecumJurisprudenciaRepository) GetGroupedByNomeCodigo() ([]model.VadeMecumJurisprudenciaGroup, error) {
	type row struct {
		Cabecalho  string `gorm:"column:cabecalho"`
		Quantidade int64  `gorm:"column:quantidade"`
	}

	var rows []row
	if err := r.db.
		Table((model.VadeMecumJurisprudencia{}).TableName()).
		Select("\"Cabecalho\" AS cabecalho, COUNT(*) AS quantidade").
		Group("\"Cabecalho\"").
		Order("\"Cabecalho\" ASC").
		Scan(&rows).Error; err != nil {
		return nil, err
	}

	groups := make([]model.VadeMecumJurisprudenciaGroup, 0, len(rows))
	for _, item := range rows {
		group := model.VadeMecumJurisprudenciaGroup{
			Cabecalho:  item.Cabecalho,
			Quantidade: item.Quantidade,
		}
		groups = append(groups, group)
	}

	return groups, nil
}

func (r *VadeMecumJurisprudenciaRepository) DeleteAll() error {
	return r.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.VadeMecumJurisprudencia{}).Error
}

func (r *VadeMecumJurisprudenciaRepository) Create(item *model.VadeMecumJurisprudencia) error {
	return r.db.Create(item).Error
}

func (r *VadeMecumJurisprudenciaRepository) GetByID(id string) (*model.VadeMecumJurisprudencia, error) {
	var item model.VadeMecumJurisprudencia
	if err := r.db.First(&item, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *VadeMecumJurisprudenciaRepository) Update(item *model.VadeMecumJurisprudencia) error {
	return r.db.Save(item).Error
}

func (r *VadeMecumJurisprudenciaRepository) Delete(id string) error {
	return r.db.Delete(&model.VadeMecumJurisprudencia{}, "id = ?", id).Error
}
