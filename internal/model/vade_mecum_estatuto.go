package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VadeMecumEstatuto struct {
	ID             uuid.UUID      `gorm:"type:uuid;primaryKey;column:id" json:"id"`
	IDTipo         string         `gorm:"column:idtipo;type:text" json:"idtipo"`
	Tipo           string         `gorm:"column:tipo;type:text" json:"tipo"`
	IDCodigo       string         `gorm:"column:idcodigo;type:text" json:"idcodigo"`
	NomeCodigo     string         `gorm:"column:nomecodigo;type:text" json:"nomecodigo"`
	Cabecalho      string         `gorm:"column:Cabecalho;type:text" json:"Cabecalho"`
	Parte          string         `gorm:"column:PARTE;type:text" json:"PARTE"`
	LivroID        string         `gorm:"column:idlivro;type:text" json:"idlivro"`
	Livro          string         `gorm:"column:livro;type:text" json:"livro"`
	LivroTexto     string         `gorm:"column:livrotexto;type:text" json:"livrotexto"`
	TituloID       string         `gorm:"column:idtitulo;type:text" json:"idtitulo"`
	Titulo         string         `gorm:"column:titulo;type:text" json:"titulo"`
	TituloTexto    string         `gorm:"column:titulotexto;type:text" json:"titulotexto"`
	SubtituloID    string         `gorm:"column:idsubtitulo;type:text" json:"idsubtitulo"`
	Subtitulo      string         `gorm:"column:subtitulo;type:text" json:"subtitulo"`
	SubtituloTexto string         `gorm:"column:subtitulotexto;type:text" json:"subtitulotexto"`
	CapituloID     string         `gorm:"column:idcapitulo;type:text" json:"idcapitulo"`
	Capitulo       string         `gorm:"column:capitulo;type:text" json:"capitulo"`
	CapituloTexto  string         `gorm:"column:capitulotexto;type:text" json:"capitulotexto"`
	SecaoID        string         `gorm:"column:idsecao;type:text" json:"idsecao"`
	Secao          string         `gorm:"column:secao;type:text" json:"secao"`
	SecaoTexto     string         `gorm:"column:secaotexto;type:text" json:"secaotexto"`
	SubsecaoID     string         `gorm:"column:idsubsecao;type:text" json:"idsubsecao"`
	Subsecao       string         `gorm:"column:subsecao;type:text" json:"subsecao"`
	SubsecaoTexto  string         `gorm:"column:subsecaotexto;type:text" json:"subsecaotexto"`
	NumeroArtigo   string         `gorm:"column:num_artigo;type:text" json:"num_artigo"`
	Artigos        string         `gorm:"column:Artigos;type:text" json:"Artigos"`
	Ordem          string         `gorm:"column:Ordem;type:text" json:"Ordem"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}

func (VadeMecumEstatuto) TableName() string {
	return "vade_mecum_estatutos"
}

func (v *VadeMecumEstatuto) BeforeCreate(tx *gorm.DB) error {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return nil
}
