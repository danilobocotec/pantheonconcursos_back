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
	query := r.db.Model(&model.Questao{})
	if filters != nil {
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
	}

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
