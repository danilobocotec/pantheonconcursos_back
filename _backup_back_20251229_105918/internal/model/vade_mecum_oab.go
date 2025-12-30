package model

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VadeMecumOAB struct {
	ID            string         `gorm:"type:text;primaryKey;column:id" json:"id"`
	IDTipo        string         `gorm:"column:idtipo;type:text" json:"idtipo"`
	Tipo          string         `gorm:"column:tipo;type:text" json:"tipo"`
	NomeCodigo    string         `gorm:"column:nomecodigo;type:text" json:"nomecodigo"`
	Cabecalho     string         `gorm:"column:Cabecalho;type:text" json:"Cabecalho"`
	Titulo        string         `gorm:"column:titulo;type:text" json:"titulo"`
	TituloTexto   string         `gorm:"column:titulotexto;type:text" json:"titulotexto"`
	TituloLabel   string         `gorm:"column:titulo_label;type:text" json:"titulo_label"`
	Capitulo      string         `gorm:"column:capitulo;type:text" json:"capitulo"`
	CapituloTexto string         `gorm:"column:capitulotexto;type:text" json:"capitulotexto"`
	CapituloLabel string         `gorm:"column:capitulo_label;type:text" json:"capitulo_label"`
	Secao         string         `gorm:"column:secao;type:text" json:"secao"`
	SecaoTexto    string         `gorm:"column:secaotexto;type:text" json:"secaotexto"`
	SecaoLabel    string         `gorm:"column:secao_label;type:text" json:"secao_label"`
	Subsecao      string         `gorm:"column:subsecao;type:text" json:"subsecao"`
	SubsecaoTexto string         `gorm:"column:subsecaotexto;type:text" json:"subsecaotexto"`
	SubsecaoLabel string         `gorm:"column:subsecao_label;type:text" json:"subsecao_label"`
	NumeroArtigo  string         `gorm:"column:num_artigo;type:text" json:"num_artigo"`
	Artigos       string         `gorm:"column:Artigos;type:text" json:"Artigos"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

func (v *VadeMecumOAB) BeforeCreate(tx *gorm.DB) error {
	if strings.TrimSpace(v.ID) == "" {
		v.ID = uuid.NewString()
	}
	return nil
}

func (VadeMecumOAB) TableName() string {
	return "vade_mecum_oab"
}

type CreateVadeMecumOABRequest struct {
	ID            string `json:"id"`
	IDTipo        string `json:"idtipo"`
	Tipo          string `json:"tipo"`
	NomeCodigo    string `json:"nomecodigo" binding:"required"`
	Cabecalho     string `json:"Cabecalho"`
	Titulo        string `json:"titulo"`
	TituloTexto   string `json:"titulotexto"`
	TituloLabel   string `json:"titulo_label"`
	Capitulo      string `json:"capitulo"`
	CapituloTexto string `json:"capitulotexto"`
	CapituloLabel string `json:"capitulo_label"`
	Secao         string `json:"secao"`
	SecaoTexto    string `json:"secaotexto"`
	SecaoLabel    string `json:"secao_label"`
	Subsecao      string `json:"subsecao"`
	SubsecaoTexto string `json:"subsecaotexto"`
	SubsecaoLabel string `json:"subsecao_label"`
	NumeroArtigo  string `json:"num_artigo"`
	Artigos       string `json:"Artigos"`
}

type UpdateVadeMecumOABRequest struct {
	IDTipo        *string `json:"idtipo"`
	Tipo          *string `json:"tipo"`
	NomeCodigo    *string `json:"nomecodigo"`
	Cabecalho     *string `json:"Cabecalho"`
	Titulo        *string `json:"titulo"`
	TituloTexto   *string `json:"titulotexto"`
	TituloLabel   *string `json:"titulo_label"`
	Capitulo      *string `json:"capitulo"`
	CapituloTexto *string `json:"capitulotexto"`
	CapituloLabel *string `json:"capitulo_label"`
	Secao         *string `json:"secao"`
	SecaoTexto    *string `json:"secaotexto"`
	SecaoLabel    *string `json:"secao_label"`
	Subsecao      *string `json:"subsecao"`
	SubsecaoTexto *string `json:"subsecaotexto"`
	SubsecaoLabel *string `json:"subsecao_label"`
	NumeroArtigo  *string `json:"num_artigo"`
	Artigos       *string `json:"Artigos"`
}
