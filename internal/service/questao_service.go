package service

import (
	"errors"
	"strings"

	"github.com/thepantheon/api/internal/model"
	"github.com/thepantheon/api/internal/repository"
)

type QuestaoService struct {
	repo *repository.QuestaoRepository
}

func NewQuestaoService(repo *repository.QuestaoRepository) *QuestaoService {
	return &QuestaoService{repo: repo}
}

func (s *QuestaoService) GetAll(filters *model.QuestaoFilters) ([]model.Questao, error) {
	return s.repo.GetAll(filters)
}

func (s *QuestaoService) GetByID(id int) (*model.Questao, error) {
	if id <= 0 {
		return nil, errors.New("id invalido")
	}
	return s.repo.GetByID(id)
}

func (s *QuestaoService) Count(filters *model.QuestaoFilters) (int64, error) {
	return s.repo.Count(filters)
}

func (s *QuestaoService) Create(req *model.CreateQuestaoRequest) (*model.Questao, error) {
	if req == nil {
		return nil, errors.New("payload obrigatorio")
	}

	item := &model.Questao{
		QuestaoID:                req.QuestaoID,
		URL:                      trimStringPtr(req.URL),
		Titulo:                   trimStringPtr(req.Titulo),
		Enunciado:                trimStringPtr(req.Enunciado),
		AlternativaA:             trimStringPtr(req.AlternativaA),
		AlternativaB:             trimStringPtr(req.AlternativaB),
		AlternativaC:             trimStringPtr(req.AlternativaC),
		AlternativaD:             trimStringPtr(req.AlternativaD),
		AlternativaE:             trimStringPtr(req.AlternativaE),
		RespostaCorreta:          trimStringPtr(req.RespostaCorreta),
		Disciplina:               trimStringPtr(req.Disciplina),
		Assunto:                  trimStringPtr(req.Assunto),
		Banca:                    trimStringPtr(req.Banca),
		Ano:                      req.Ano,
		Nivel:                    trimStringPtr(req.Nivel),
		HTMLCompleto:             trimStringPtr(req.HTMLCompleto),
		DataCaptura:              req.DataCaptura,
		DataAtualizacao:          req.DataAtualizacao,
		TipoQuestao:              trimStringPtr(req.TipoQuestao),
		CorrecaoQuestao:          req.CorrecaoQuestao,
		NumeroAlternativaCorreta: req.NumeroAlternativaCorreta,
		Anulada:                  req.Anulada,
		Desatualizada:            req.Desatualizada,
		PossuiResolucaoBanca:     req.PossuiResolucaoBanca,
		GabaritoPreliminar:       req.GabaritoPreliminar,
		QuestaoOculta:            req.QuestaoOculta,
		IDQuestao:                trimStringPtr(req.IDQuestao),
		IDQuestaoOriginal:        trimStringPtr(req.IDQuestaoOriginal),
		Gabarito:                 trimStringPtr(req.Gabarito),
		Comentario:               trimStringPtr(req.Comentario),
		ResolucaoBanca:           trimStringPtr(req.ResolucaoBanca),
		Instituicao:              trimStringPtr(req.Instituicao),
		Cargo:                    trimStringPtr(req.Cargo),
		Orgao:                    trimStringPtr(req.Orgao),
		Localizacao:              trimStringPtr(req.Localizacao),
		Dificuldade:              trimStringPtr(req.Dificuldade),
		QuantidadeResolucoes:     req.QuantidadeResolucoes,
		AcertosPercentual:        req.AcertosPercentual,
		AreaConhecimento:         trimStringPtr(req.AreaConhecimento),
		Concurso:                 trimStringPtr(req.Concurso),
		FormatoQuestao:           trimStringPtr(req.FormatoQuestao),
		TipoProva:                trimStringPtr(req.TipoProva),
	}

	if req.CamposJSON != nil {
		item.CamposJSON = *req.CamposJSON
	}

	if err := s.repo.Create(item); err != nil {
		return nil, err
	}

	return item, nil
}

func (s *QuestaoService) Update(id int, req *model.UpdateQuestaoRequest) (*model.Questao, error) {
	if req == nil {
		return nil, errors.New("payload obrigatorio")
	}
	if id <= 0 {
		return nil, errors.New("id invalido")
	}

	item, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.QuestaoID != nil {
		item.QuestaoID = req.QuestaoID
	}
	if req.URL != nil {
		item.URL = trimStringPtr(req.URL)
	}
	if req.Titulo != nil {
		item.Titulo = trimStringPtr(req.Titulo)
	}
	if req.Enunciado != nil {
		item.Enunciado = trimStringPtr(req.Enunciado)
	}
	if req.AlternativaA != nil {
		item.AlternativaA = trimStringPtr(req.AlternativaA)
	}
	if req.AlternativaB != nil {
		item.AlternativaB = trimStringPtr(req.AlternativaB)
	}
	if req.AlternativaC != nil {
		item.AlternativaC = trimStringPtr(req.AlternativaC)
	}
	if req.AlternativaD != nil {
		item.AlternativaD = trimStringPtr(req.AlternativaD)
	}
	if req.AlternativaE != nil {
		item.AlternativaE = trimStringPtr(req.AlternativaE)
	}
	if req.RespostaCorreta != nil {
		item.RespostaCorreta = trimStringPtr(req.RespostaCorreta)
	}
	if req.Disciplina != nil {
		item.Disciplina = trimStringPtr(req.Disciplina)
	}
	if req.Assunto != nil {
		item.Assunto = trimStringPtr(req.Assunto)
	}
	if req.Banca != nil {
		item.Banca = trimStringPtr(req.Banca)
	}
	if req.Ano != nil {
		item.Ano = req.Ano
	}
	if req.Nivel != nil {
		item.Nivel = trimStringPtr(req.Nivel)
	}
	if req.CamposJSON != nil {
		item.CamposJSON = *req.CamposJSON
	}
	if req.HTMLCompleto != nil {
		item.HTMLCompleto = trimStringPtr(req.HTMLCompleto)
	}
	if req.DataCaptura != nil {
		item.DataCaptura = req.DataCaptura
	}
	if req.DataAtualizacao != nil {
		item.DataAtualizacao = req.DataAtualizacao
	}
	if req.TipoQuestao != nil {
		item.TipoQuestao = trimStringPtr(req.TipoQuestao)
	}
	if req.CorrecaoQuestao != nil {
		item.CorrecaoQuestao = req.CorrecaoQuestao
	}
	if req.NumeroAlternativaCorreta != nil {
		item.NumeroAlternativaCorreta = req.NumeroAlternativaCorreta
	}
	if req.Anulada != nil {
		item.Anulada = req.Anulada
	}
	if req.Desatualizada != nil {
		item.Desatualizada = req.Desatualizada
	}
	if req.PossuiResolucaoBanca != nil {
		item.PossuiResolucaoBanca = req.PossuiResolucaoBanca
	}
	if req.GabaritoPreliminar != nil {
		item.GabaritoPreliminar = req.GabaritoPreliminar
	}
	if req.QuestaoOculta != nil {
		item.QuestaoOculta = req.QuestaoOculta
	}
	if req.IDQuestao != nil {
		item.IDQuestao = trimStringPtr(req.IDQuestao)
	}
	if req.IDQuestaoOriginal != nil {
		item.IDQuestaoOriginal = trimStringPtr(req.IDQuestaoOriginal)
	}
	if req.Gabarito != nil {
		item.Gabarito = trimStringPtr(req.Gabarito)
	}
	if req.Comentario != nil {
		item.Comentario = trimStringPtr(req.Comentario)
	}
	if req.ResolucaoBanca != nil {
		item.ResolucaoBanca = trimStringPtr(req.ResolucaoBanca)
	}
	if req.Instituicao != nil {
		item.Instituicao = trimStringPtr(req.Instituicao)
	}
	if req.Cargo != nil {
		item.Cargo = trimStringPtr(req.Cargo)
	}
	if req.Orgao != nil {
		item.Orgao = trimStringPtr(req.Orgao)
	}
	if req.Localizacao != nil {
		item.Localizacao = trimStringPtr(req.Localizacao)
	}
	if req.Dificuldade != nil {
		item.Dificuldade = trimStringPtr(req.Dificuldade)
	}
	if req.QuantidadeResolucoes != nil {
		item.QuantidadeResolucoes = req.QuantidadeResolucoes
	}
	if req.AcertosPercentual != nil {
		item.AcertosPercentual = req.AcertosPercentual
	}
	if req.AreaConhecimento != nil {
		item.AreaConhecimento = trimStringPtr(req.AreaConhecimento)
	}
	if req.Concurso != nil {
		item.Concurso = trimStringPtr(req.Concurso)
	}
	if req.FormatoQuestao != nil {
		item.FormatoQuestao = trimStringPtr(req.FormatoQuestao)
	}
	if req.TipoProva != nil {
		item.TipoProva = trimStringPtr(req.TipoProva)
	}

	if err := s.repo.Update(item); err != nil {
		return nil, err
	}

	return item, nil
}

func (s *QuestaoService) Delete(id int) error {
	if id <= 0 {
		return errors.New("id invalido")
	}
	return s.repo.Delete(id)
}

func (s *QuestaoService) GetFilterOptions() (*model.QuestaoFiltersResponse, error) {
	return s.repo.GetFilterOptions()
}

func trimStringPtr(value *string) *string {
	if value == nil {
		return nil
	}
	trimmed := strings.TrimSpace(*value)
	return &trimmed
}
