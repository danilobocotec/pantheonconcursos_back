package repository

import (
	"github.com/thepantheon/api/internal/model"
	"gorm.io/gorm"
)

type QuestaoRepository struct {
	db *gorm.DB
}

func NewQuestaoRepository(db *gorm.DB) *QuestaoRepository {
	return &QuestaoRepository{db: db}
}

func (r *QuestaoRepository) Create(item *model.Questao) error {
	return r.db.Create(item).Error
}

func (r *QuestaoRepository) GetAll(filters *model.QuestaoFilters) ([]model.Questao, error) {
	var items []model.Questao
	query := r.buildQuestaoQuery(filters)

	if err := query.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *QuestaoRepository) GetByID(id int) (*model.Questao, error) {
	var item model.Questao
	if err := r.db.First(&item, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *QuestaoRepository) Update(item *model.Questao) error {
	return r.db.Save(item).Error
}

func (r *QuestaoRepository) Delete(id int) error {
	return r.db.Delete(&model.Questao{}, "id = ?", id).Error
}

func (r *QuestaoRepository) Count(filters *model.QuestaoFilters) (int64, error) {
	var count int64
	err := r.buildQuestaoQuery(filters).
		Where("erro_captura IS NULL OR erro_captura = ?", false).
		Count(&count).Error
	return count, err
}

func (r *QuestaoRepository) GetFilterOptions() (*model.QuestaoFiltersResponse, error) {
	response := &model.QuestaoFiltersResponse{}

	loadDistinct := func(column string, dest *[]string) error {
		return r.db.
			Model(&model.Questao{}).
			Where(column+" IS NOT NULL AND "+column+" <> ''").
			Distinct().
			Order(column).
			Pluck(column, dest).Error
	}

	if err := loadDistinct("disciplina", &response.Disciplina); err != nil {
		return nil, err
	}
	if err := loadDistinct("assunto", &response.Assunto); err != nil {
		return nil, err
	}
	if err := loadDistinct("banca", &response.Banca); err != nil {
		return nil, err
	}
	if err := loadDistinct("orgao", &response.Orgao); err != nil {
		return nil, err
	}
	if err := loadDistinct("cargo", &response.Cargo); err != nil {
		return nil, err
	}
	if err := loadDistinct("concurso", &response.Concurso); err != nil {
		return nil, err
	}
	if err := loadDistinct("area_conhecimento", &response.AreaConhecimento); err != nil {
		return nil, err
	}

	return response, nil
}

func (r *QuestaoRepository) buildQuestaoQuery(filters *model.QuestaoFilters) *gorm.DB {
	query := r.db.Model(&model.Questao{})
	if filters == nil {
		return query
	}
	if filters.Disciplina != nil {
		query = query.Where("disciplina = ?", *filters.Disciplina)
	}
	if filters.Assunto != nil {
		query = query.Where("assunto = ?", *filters.Assunto)
	}
	if filters.Banca != nil {
		query = query.Where("banca = ?", *filters.Banca)
	}
	if filters.Orgao != nil {
		query = query.Where("orgao = ?", *filters.Orgao)
	}
	if filters.Cargo != nil {
		query = query.Where("cargo = ?", *filters.Cargo)
	}
	if filters.Concurso != nil {
		query = query.Where("concurso = ?", *filters.Concurso)
	}
	if filters.AreaConhecimento != nil {
		query = query.Where("area_conhecimento = ?", *filters.AreaConhecimento)
	}
	return query
}
