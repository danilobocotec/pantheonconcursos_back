package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VadeMecumCodigo struct {
	ID             uuid.UUID      `gorm:"type:uuid;primaryKey;column:id" json:"id"`
	IDTipo         string         `gorm:"column:idtipo;type:varchar" json:"idtipo"`
	Tipo           string         `gorm:"column:tipo;type:varchar" json:"tipo"`
	IDCodigo       string         `gorm:"column:idcodigo;type:varchar" json:"idcodigo"`
	NomeCodigo     string         `gorm:"not null;column:nomecodigo;type:varchar" json:"nomecodigo"`
	Cabecalho      string         `gorm:"column:Cabecalho;type:varchar" json:"Cabecalho"`
	Parte          string         `gorm:"column:PARTE;type:varchar" json:"PARTE"`
	LivroID        string         `gorm:"column:idlivro;type:varchar" json:"idlivro"`
	Livro          string         `gorm:"column:livro;type:varchar" json:"livro"`
	LivroTexto     string         `gorm:"column:livrotexto;type:varchar" json:"livrotexto"`
	TituloID       string         `gorm:"column:idtitulo;type:varchar" json:"idtitulo"`
	Titulo         string         `gorm:"column:titulo;type:varchar" json:"titulo"`
	TituloTexto    string         `gorm:"column:titulotexto;type:varchar" json:"titulotexto"`
	SubtituloID    string         `gorm:"column:idsubtitulo;type:varchar" json:"idsubtitulo"`
	Subtitulo      string         `gorm:"column:subtitulo;type:varchar" json:"subtitulo"`
	SubtituloTexto string         `gorm:"column:subtitulotexto;type:varchar" json:"subtitulotexto"`
	CapituloID     string         `gorm:"column:idcapitulo;type:varchar" json:"idcapitulo"`
	Capitulo       string         `gorm:"column:capitulo;type:varchar" json:"capitulo"`
	CapituloTexto  string         `gorm:"column:capitulotexto;type:varchar" json:"capitulotexto"`
	SecaoID        string         `gorm:"column:idsecao;type:varchar" json:"idsecao"`
	Secao          string         `gorm:"column:secao;type:varchar" json:"secao"`
	SecaoTexto     string         `gorm:"column:secaotexto;type:varchar" json:"secaotexto"`
	SubsecaoID     string         `gorm:"column:idsubsecao;type:varchar" json:"idsubsecao"`
	Subsecao       string         `gorm:"column:subsecao;type:varchar" json:"subsecao"`
	SubsecaoTexto  string         `gorm:"column:subsecaotexto;type:varchar" json:"subsecaotexto"`
	NumeroArtigo   string         `gorm:"column:num_artigo;type:varchar" json:"num_artigo"`
	Normativo      string         `gorm:"column:Normativo;type:varchar" json:"Normativo"`
	Ordem          string         `gorm:"column:Ordem;type:varchar" json:"Ordem"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}

func (v *VadeMecumCodigo) BeforeCreate(tx *gorm.DB) error {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return nil
}

type CreateVadeMecumCodigoRequest struct {
	IDTipo         string `json:"idtipo"`
	Tipo           string `json:"tipo"`
	IDCodigo       string `json:"idcodigo"`
	NomeCodigo     string `json:"nomecodigo" binding:"required"`
	Cabecalho      string `json:"Cabecalho"`
	Parte          string `json:"PARTE"`
	LivroID        string `json:"idlivro"`
	Livro          string `json:"livro"`
	LivroTexto     string `json:"livrotexto"`
	TituloID       string `json:"idtitulo"`
	Titulo         string `json:"titulo"`
	TituloTexto    string `json:"titulotexto"`
	SubtituloID    string `json:"idsubtitulo"`
	Subtitulo      string `json:"subtitulo"`
	SubtituloTexto string `json:"subtitulotexto"`
	CapituloID     string `json:"idcapitulo"`
	Capitulo       string `json:"capitulo"`
	CapituloTexto  string `json:"capitulotexto"`
	SecaoID        string `json:"idsecao"`
	Secao          string `json:"secao"`
	SecaoTexto     string `json:"secaotexto"`
	SubsecaoID     string `json:"idsubsecao"`
	Subsecao       string `json:"subsecao"`
	SubsecaoTexto  string `json:"subsecaotexto"`
	NumeroArtigo   string `json:"num_artigo"`
	Normativo      string `json:"Normativo"`
	Ordem          string `json:"Ordem"`
}

type UpdateVadeMecumCodigoRequest struct {
	IDTipo         string `json:"idtipo"`
	Tipo           string `json:"tipo"`
	IDCodigo       string `json:"idcodigo"`
	NomeCodigo     string `json:"nomecodigo"`
	Cabecalho      string `json:"Cabecalho"`
	Parte          string `json:"PARTE"`
	LivroID        string `json:"idlivro"`
	Livro          string `json:"livro"`
	LivroTexto     string `json:"livrotexto"`
	TituloID       string `json:"idtitulo"`
	Titulo         string `json:"titulo"`
	TituloTexto    string `json:"titulotexto"`
	SubtituloID    string `json:"idsubtitulo"`
	Subtitulo      string `json:"subtitulo"`
	SubtituloTexto string `json:"subtitulotexto"`
	CapituloID     string `json:"idcapitulo"`
	Capitulo       string `json:"capitulo"`
	CapituloTexto  string `json:"capitulotexto"`
	SecaoID        string `json:"idsecao"`
	Secao          string `json:"secao"`
	SecaoTexto     string `json:"secaotexto"`
	SubsecaoID     string `json:"idsubsecao"`
	Subsecao       string `json:"subsecao"`
	SubsecaoTexto  string `json:"subsecaotexto"`
	NumeroArtigo   string `json:"num_artigo"`
	Normativo      string `json:"Normativo"`
	Ordem          string `json:"Ordem"`
}
