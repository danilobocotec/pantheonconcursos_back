package model

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// VadeMecumJurisprudencia representa registros de jurisprudência do vade-mécum.
type VadeMecumJurisprudencia struct {
	ID            string         `gorm:"type:text;primaryKey;column:id" json:"id"`
	IDTipo        string         `gorm:"column:idtipo;type:text" json:"idtipo"`
	Tipo          string         `gorm:"column:tipo;type:text" json:"tipo"`
	IDCodigo      string         `gorm:"column:idcodigo;type:text" json:"idcodigo"`
	NomeCodigo    string         `gorm:"column:nomecodigo;type:text" json:"nomecodigo"`
	Cabecalho     string         `gorm:"column:Cabecalho;type:text" json:"Cabecalho"`
	TipoDescricao string         `gorm:"column:Tipo;type:text" json:"Tipo"`
	IDRamo        string         `gorm:"column:idramo;type:text" json:"idramo"`
	RamoTexto     string         `gorm:"column:ramotexto;type:text" json:"ramotexto"`
	IDAssunto     string         `gorm:"column:idassunto;type:text" json:"idassunto"`
	AssuntoTexto  string         `gorm:"column:assuntotexto;type:text" json:"assuntotexto"`
	IDEnunciado   string         `gorm:"column:idenunciado;type:text" json:"idenunciado"`
	Enunciado     string         `gorm:"column:Enunciado;type:text" json:"Enunciado"`
	IDSecao       string         `gorm:"column:idsecao;type:text" json:"idsecao"`
	Secao         string         `gorm:"column:secao;type:text" json:"secao"`
	SecaoTexto    string         `gorm:"column:secaotexto;type:text" json:"secaotexto"`
	IDSubsecao    string         `gorm:"column:idsubsecao;type:text" json:"idsubsecao"`
	Subsecao      string         `gorm:"column:subsecao;type:text" json:"subsecao"`
	SubsecaoTexto string         `gorm:"column:subsecaotexto;type:text" json:"subsecaotexto"`
	NumeroArtigo  string         `gorm:"column:num_artigo;type:text" json:"num_artigo"`
	Normativo     string         `gorm:"column:Normativo;type:text" json:"Normativo"`
	Ordem         string         `gorm:"column:Ordem;type:text" json:"Ordem"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

func (VadeMecumJurisprudencia) TableName() string {
	return "vade_mecum_jurisprudencia"
}

func (v *VadeMecumJurisprudencia) BeforeCreate(tx *gorm.DB) error {
	if strings.TrimSpace(v.ID) == "" {
		v.ID = uuid.NewString()
	}
	return nil
}

type CreateVadeMecumJurisprudenciaRequest struct {
	ID            string `json:"id"`
	IDTipo        string `json:"idtipo"`
	Tipo          string `json:"tipo"`
	IDCodigo      string `json:"idcodigo"`
	NomeCodigo    string `json:"nomecodigo" binding:"required"`
	Cabecalho     string `json:"Cabecalho"`
	TipoDescricao string `json:"Tipo"`
	IDRamo        string `json:"idramo"`
	RamoTexto     string `json:"ramotexto"`
	IDAssunto     string `json:"idassunto"`
	AssuntoTexto  string `json:"assuntotexto"`
	IDEnunciado   string `json:"idenunciado"`
	Enunciado     string `json:"Enunciado"`
	IDSecao       string `json:"idsecao"`
	Secao         string `json:"secao"`
	SecaoTexto    string `json:"secaotexto"`
	IDSubsecao    string `json:"idsubsecao"`
	Subsecao      string `json:"subsecao"`
	SubsecaoTexto string `json:"subsecaotexto"`
	NumeroArtigo  string `json:"num_artigo"`
	Normativo     string `json:"Normativo"`
	Ordem         string `json:"Ordem"`
}

type UpdateVadeMecumJurisprudenciaRequest struct {
	IDTipo        *string `json:"idtipo"`
	Tipo          *string `json:"tipo"`
	IDCodigo      *string `json:"idcodigo"`
	NomeCodigo    *string `json:"nomecodigo"`
	Cabecalho     *string `json:"Cabecalho"`
	TipoDescricao *string `json:"Tipo"`
	IDRamo        *string `json:"idramo"`
	RamoTexto     *string `json:"ramotexto"`
	IDAssunto     *string `json:"idassunto"`
	AssuntoTexto  *string `json:"assuntotexto"`
	IDEnunciado   *string `json:"idenunciado"`
	Enunciado     *string `json:"Enunciado"`
	IDSecao       *string `json:"idsecao"`
	Secao         *string `json:"secao"`
	SecaoTexto    *string `json:"secaotexto"`
	IDSubsecao    *string `json:"idsubsecao"`
	Subsecao      *string `json:"subsecao"`
	SubsecaoTexto *string `json:"subsecaotexto"`
	NumeroArtigo  *string `json:"num_artigo"`
	Normativo     *string `json:"Normativo"`
	Ordem         *string `json:"Ordem"`
}
