package model

import (
	"time"

	"gorm.io/datatypes"
)

type Questao struct {
	ID                       int             `gorm:"primaryKey;column:id" json:"id"`
	QuestaoID                *int            `gorm:"column:questao_id" json:"questao_id"`
	URL                      *string         `gorm:"column:url;type:varchar(500)" json:"url"`
	Titulo                   *string         `gorm:"column:titulo;type:text" json:"titulo"`
	Enunciado                *string         `gorm:"column:enunciado;type:text" json:"enunciado"`
	AlternativaA             *string         `gorm:"column:alternativa_a;type:text" json:"alternativa_a"`
	AlternativaB             *string         `gorm:"column:alternativa_b;type:text" json:"alternativa_b"`
	AlternativaC             *string         `gorm:"column:alternativa_c;type:text" json:"alternativa_c"`
	AlternativaD             *string         `gorm:"column:alternativa_d;type:text" json:"alternativa_d"`
	AlternativaE             *string         `gorm:"column:alternativa_e;type:text" json:"alternativa_e"`
	RespostaCorreta          *string         `gorm:"column:resposta_correta;type:text" json:"resposta_correta"`
	Disciplina               *string         `gorm:"column:disciplina;type:varchar(200)" json:"disciplina"`
	Assunto                  *string         `gorm:"column:assunto;type:varchar(200)" json:"assunto"`
	Banca                    *string         `gorm:"column:banca;type:varchar(200)" json:"banca"`
	Ano                      *int            `gorm:"column:ano" json:"ano"`
	Nivel                    *string         `gorm:"column:nivel;type:varchar(100)" json:"nivel"`
	CamposJSON               datatypes.JSON  `gorm:"column:campos_json;type:jsonb" json:"campos_json"`
	HTMLCompleto             *string         `gorm:"column:html_completo;type:text" json:"html_completo"`
	DataCaptura              *time.Time      `gorm:"column:data_captura" json:"data_captura"`
	DataAtualizacao          *time.Time      `gorm:"column:data_atualizacao" json:"data_atualizacao"`
	TipoQuestao              *string         `gorm:"column:tipo_questao;type:varchar(50)" json:"tipo_questao"`
	CorrecaoQuestao          *bool           `gorm:"column:correcao_questao" json:"correcao_questao"`
	NumeroAlternativaCorreta *int            `gorm:"column:numero_alternativa_correta" json:"numero_alternativa_correta"`
	Anulada                  *bool           `gorm:"column:anulada" json:"anulada"`
	Desatualizada            *bool           `gorm:"column:desatualizada" json:"desatualizada"`
	PossuiResolucaoBanca     *bool           `gorm:"column:possui_resolucao_banca" json:"possui_resolucao_banca"`
	GabaritoPreliminar       *bool           `gorm:"column:gabarito_preliminar" json:"gabarito_preliminar"`
	QuestaoOculta            *bool           `gorm:"column:questao_oculta" json:"questao_oculta"`
	IDQuestao                *string         `gorm:"column:id_questao;type:varchar(100)" json:"id_questao"`
	IDQuestaoOriginal        *string         `gorm:"column:id_questao_original;type:varchar(100)" json:"id_questao_original"`
	Gabarito                 *string         `gorm:"column:gabarito;type:text" json:"gabarito"`
	Comentario               *string         `gorm:"column:comentario;type:text" json:"comentario"`
	ResolucaoBanca           *string         `gorm:"column:resolucao_banca;type:text" json:"resolucao_banca"`
	Instituicao              *string         `gorm:"column:instituicao;type:varchar(200)" json:"instituicao"`
	Cargo                    *string         `gorm:"column:cargo;type:varchar(200)" json:"cargo"`
	Orgao                    *string         `gorm:"column:orgao;type:varchar(200)" json:"orgao"`
	Localizacao              *string         `gorm:"column:localizacao;type:varchar(200)" json:"localizacao"`
	Dificuldade              *string         `gorm:"column:dificuldade;type:varchar(50)" json:"dificuldade"`
	QuantidadeResolucoes     *int            `gorm:"column:quantidade_resolucoes" json:"quantidade_resolucoes"`
	AcertosPercentual        *float64        `gorm:"column:acertos_percentual" json:"acertos_percentual"`
	AreaConhecimento         *string         `gorm:"column:area_conhecimento;type:varchar(200)" json:"area_conhecimento"`
	Concurso                 *string         `gorm:"column:concurso;type:varchar(200)" json:"concurso"`
	FormatoQuestao           *string         `gorm:"column:formato_questao;type:varchar(100)" json:"formato_questao"`
	TipoProva                *string         `gorm:"column:tipo_prova;type:varchar(100)" json:"tipo_prova"`
}

func (Questao) TableName() string {
	return "questoes"
}

type CreateQuestaoRequest struct {
	QuestaoID                *int            `json:"questao_id"`
	URL                      *string         `json:"url"`
	Titulo                   *string         `json:"titulo"`
	Enunciado                *string         `json:"enunciado"`
	AlternativaA             *string         `json:"alternativa_a"`
	AlternativaB             *string         `json:"alternativa_b"`
	AlternativaC             *string         `json:"alternativa_c"`
	AlternativaD             *string         `json:"alternativa_d"`
	AlternativaE             *string         `json:"alternativa_e"`
	RespostaCorreta          *string         `json:"resposta_correta"`
	Disciplina               *string         `json:"disciplina"`
	Assunto                  *string         `json:"assunto"`
	Banca                    *string         `json:"banca"`
	Ano                      *int            `json:"ano"`
	Nivel                    *string         `json:"nivel"`
	CamposJSON               *datatypes.JSON `json:"campos_json"`
	HTMLCompleto             *string         `json:"html_completo"`
	DataCaptura              *time.Time      `json:"data_captura"`
	DataAtualizacao          *time.Time      `json:"data_atualizacao"`
	TipoQuestao              *string         `json:"tipo_questao"`
	CorrecaoQuestao          *bool           `json:"correcao_questao"`
	NumeroAlternativaCorreta *int            `json:"numero_alternativa_correta"`
	Anulada                  *bool           `json:"anulada"`
	Desatualizada            *bool           `json:"desatualizada"`
	PossuiResolucaoBanca     *bool           `json:"possui_resolucao_banca"`
	GabaritoPreliminar       *bool           `json:"gabarito_preliminar"`
	QuestaoOculta            *bool           `json:"questao_oculta"`
	IDQuestao                *string         `json:"id_questao"`
	IDQuestaoOriginal        *string         `json:"id_questao_original"`
	Gabarito                 *string         `json:"gabarito"`
	Comentario               *string         `json:"comentario"`
	ResolucaoBanca           *string         `json:"resolucao_banca"`
	Instituicao              *string         `json:"instituicao"`
	Cargo                    *string         `json:"cargo"`
	Orgao                    *string         `json:"orgao"`
	Localizacao              *string         `json:"localizacao"`
	Dificuldade              *string         `json:"dificuldade"`
	QuantidadeResolucoes     *int            `json:"quantidade_resolucoes"`
	AcertosPercentual        *float64        `json:"acertos_percentual"`
	AreaConhecimento         *string         `json:"area_conhecimento"`
	Concurso                 *string         `json:"concurso"`
	FormatoQuestao           *string         `json:"formato_questao"`
	TipoProva                *string         `json:"tipo_prova"`
}

type UpdateQuestaoRequest struct {
	QuestaoID                *int            `json:"questao_id"`
	URL                      *string         `json:"url"`
	Titulo                   *string         `json:"titulo"`
	Enunciado                *string         `json:"enunciado"`
	AlternativaA             *string         `json:"alternativa_a"`
	AlternativaB             *string         `json:"alternativa_b"`
	AlternativaC             *string         `json:"alternativa_c"`
	AlternativaD             *string         `json:"alternativa_d"`
	AlternativaE             *string         `json:"alternativa_e"`
	RespostaCorreta          *string         `json:"resposta_correta"`
	Disciplina               *string         `json:"disciplina"`
	Assunto                  *string         `json:"assunto"`
	Banca                    *string         `json:"banca"`
	Ano                      *int            `json:"ano"`
	Nivel                    *string         `json:"nivel"`
	CamposJSON               *datatypes.JSON `json:"campos_json"`
	HTMLCompleto             *string         `json:"html_completo"`
	DataCaptura              *time.Time      `json:"data_captura"`
	DataAtualizacao          *time.Time      `json:"data_atualizacao"`
	TipoQuestao              *string         `json:"tipo_questao"`
	CorrecaoQuestao          *bool           `json:"correcao_questao"`
	NumeroAlternativaCorreta *int            `json:"numero_alternativa_correta"`
	Anulada                  *bool           `json:"anulada"`
	Desatualizada            *bool           `json:"desatualizada"`
	PossuiResolucaoBanca     *bool           `json:"possui_resolucao_banca"`
	GabaritoPreliminar       *bool           `json:"gabarito_preliminar"`
	QuestaoOculta            *bool           `json:"questao_oculta"`
	IDQuestao                *string         `json:"id_questao"`
	IDQuestaoOriginal        *string         `json:"id_questao_original"`
	Gabarito                 *string         `json:"gabarito"`
	Comentario               *string         `json:"comentario"`
	ResolucaoBanca           *string         `json:"resolucao_banca"`
	Instituicao              *string         `json:"instituicao"`
	Cargo                    *string         `json:"cargo"`
	Orgao                    *string         `json:"orgao"`
	Localizacao              *string         `json:"localizacao"`
	Dificuldade              *string         `json:"dificuldade"`
	QuantidadeResolucoes     *int            `json:"quantidade_resolucoes"`
	AcertosPercentual        *float64        `json:"acertos_percentual"`
	AreaConhecimento         *string         `json:"area_conhecimento"`
	Concurso                 *string         `json:"concurso"`
	FormatoQuestao           *string         `json:"formato_questao"`
	TipoProva                *string         `json:"tipo_prova"`
}

type QuestaoFilters struct {
	Disciplina       *string
	Assunto          *string
	Banca            *string
	Orgao            *string
	Cargo            *string
	Concurso         *string
	AreaConhecimento *string
}

type QuestaoFiltersResponse struct {
	Disciplina       []string `json:"disciplina"`
	Assunto          []string `json:"assunto"`
	Banca            []string `json:"banca"`
	Orgao            []string `json:"orgao"`
	Cargo            []string `json:"cargo"`
	Concurso         []string `json:"concurso"`
	AreaConhecimento []string `json:"area_conhecimento"`
}

type QuestaoCountResponse struct {
	Count int64 `json:"count"`
}
