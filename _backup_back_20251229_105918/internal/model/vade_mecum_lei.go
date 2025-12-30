package model

import (
	"time"

	"gorm.io/gorm"
)

type VadeMecumLei struct {
	ID            string         `gorm:"type:text;primaryKey;column:id" json:"id"`
	IDTipo        string         `gorm:"column:idtipo;type:text" json:"idtipo"`
	Tipo          string         `gorm:"column:tipo;type:text" json:"tipo"`
	NomeCodigo    string         `gorm:"column:nomecodigo;type:text" json:"nomecodigo"`
	Cabecalho     string         `gorm:"column:Cabecalho;type:text" json:"Cabecalho"`
	IDParte       string         `gorm:"column:idPARTE;type:text" json:"idPARTE"`
	Parte         string         `gorm:"column:PARTE;type:text" json:"PARTE"`
	ParteTexto    string         `gorm:"column:PARTETEXTO;type:text" json:"PARTETEXTO"`
	IDTitulo      string         `gorm:"column:idtitulo;type:text" json:"idtitulo"`
	Titulo        string         `gorm:"column:titulo;type:text" json:"titulo"`
	TituloTexto   string         `gorm:"column:titulotexto;type:text" json:"titulotexto"`
	IDCapitulo    string         `gorm:"column:idcapitulo;type:text" json:"idcapitulo"`
	Capitulo      string         `gorm:"column:capitulo;type:text" json:"capitulo"`
	CapituloTexto string         `gorm:"column:capitulotexto;type:text" json:"capitulotexto"`
	IDSecao       string         `gorm:"column:idsecao;type:text" json:"idsecao"`
	Secao         string         `gorm:"column:secao;type:text" json:"secao"`
	SecaoTexto    string         `gorm:"column:secaotexto;type:text" json:"secaotexto"`
	IDSubsecao    string         `gorm:"column:idsubsecao;type:text" json:"idsubsecao"`
	Subsecao      string         `gorm:"column:subsecao;type:text" json:"subsecao"`
	SubsecaoTexto string         `gorm:"column:subsecaotexto;type:text" json:"subsecaotexto"`
	NumeroArtigo  string         `gorm:"column:num_artigo;type:text" json:"num_artigo"`
	Artigos       string         `gorm:"column:Artigos;type:text" json:"Artigos"`
	Ordem         string         `gorm:"column:Ordem;type:text" json:"Ordem"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

func (VadeMecumLei) TableName() string {
	return "vade_mecum_leis"
}

type CreateVadeMecumLeiRequest struct {
	ID            string `json:"id"`
	IDTipo        string `json:"idtipo"`
	Tipo          string `json:"tipo"`
	NomeCodigo    string `json:"nomecodigo" binding:"required"`
	Cabecalho     string `json:"Cabecalho"`
	IDParte       string `json:"idPARTE"`
	Parte         string `json:"PARTE"`
	ParteTexto    string `json:"PARTETEXTO"`
	IDTitulo      string `json:"idtitulo"`
	Titulo        string `json:"titulo"`
	TituloTexto   string `json:"titulotexto"`
	IDCapitulo    string `json:"idcapitulo"`
	Capitulo      string `json:"capitulo"`
	CapituloTexto string `json:"capitulotexto"`
	IDSecao       string `json:"idsecao"`
	Secao         string `json:"secao"`
	SecaoTexto    string `json:"secaotexto"`
	IDSubsecao    string `json:"idsubsecao"`
	Subsecao      string `json:"subsecao"`
	SubsecaoTexto string `json:"subsecaotexto"`
	NumeroArtigo  string `json:"num_artigo"`
	Artigos       string `json:"Artigos"`
	Ordem         string `json:"Ordem"`
}

type UpdateVadeMecumLeiRequest struct {
	IDTipo        *string `json:"idtipo"`
	Tipo          *string `json:"tipo"`
	NomeCodigo    *string `json:"nomecodigo"`
	Cabecalho     *string `json:"Cabecalho"`
	IDParte       *string `json:"idPARTE"`
	Parte         *string `json:"PARTE"`
	ParteTexto    *string `json:"PARTETEXTO"`
	IDTitulo      *string `json:"idtitulo"`
	Titulo        *string `json:"titulo"`
	TituloTexto   *string `json:"titulotexto"`
	IDCapitulo    *string `json:"idcapitulo"`
	Capitulo      *string `json:"capitulo"`
	CapituloTexto *string `json:"capitulotexto"`
	IDSecao       *string `json:"idsecao"`
	Secao         *string `json:"secao"`
	SecaoTexto    *string `json:"secaotexto"`
	IDSubsecao    *string `json:"idsubsecao"`
	Subsecao      *string `json:"subsecao"`
	SubsecaoTexto *string `json:"subsecaotexto"`
	NumeroArtigo  *string `json:"num_artigo"`
	Artigos       *string `json:"Artigos"`
	Ordem         *string `json:"Ordem"`
}

type VadeMecumLeiGrupoServico struct {
	NomeCodigo string `json:"nomecodigo"`
}
