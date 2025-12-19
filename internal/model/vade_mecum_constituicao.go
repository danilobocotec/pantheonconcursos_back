package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VadeMecumConstituicao struct {
	ID            uuid.UUID      `gorm:"type:uuid;primaryKey;column:id" json:"id"`
	RegistroID    string         `gorm:"column:registro_id;type:text" json:"registro_id"`
	IDTipo        string         `gorm:"column:idtipo;type:text" json:"idtipo"`
	Tipo          string         `gorm:"column:tipo;type:text" json:"tipo"`
	Cabecalho     string         `gorm:"column:cabecalho;type:text" json:"cabecalho"`
	IDTitulo      string         `gorm:"column:idtitulo;type:text" json:"idtitulo"`
	Titulo        string         `gorm:"column:titulo;type:text" json:"titulo"`
	TextoDoTitulo string         `gorm:"column:textodotitulo;type:text" json:"textodotitulo"`
	IDCapitulo    string         `gorm:"column:idcapitulo;type:text" json:"idcapitulo"`
	Capitulo      string         `gorm:"column:capitulo;type:text" json:"capitulo"`
	TextoCapitulo string         `gorm:"column:textocapitulo;type:text" json:"textocapitulo"`
	IDSecao       string         `gorm:"column:idsecao;type:text" json:"idsecao"`
	Secao         string         `gorm:"column:secao;type:text" json:"secao"`
	TextoSecao    string         `gorm:"column:textosecao;type:text" json:"textosecao"`
	IDSubsecao    string         `gorm:"column:idsubsecao;type:text" json:"idsubsecao"`
	Subsecao      string         `gorm:"column:subsecao;type:text" json:"subsecao"`
	TextoSubsecao string         `gorm:"column:subsecaotexto;type:text" json:"subsecaotexto"`
	Normativo     string         `gorm:"column:Normativo;type:text" json:"Normativo"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

func (VadeMecumConstituicao) TableName() string {
	return "vade_mecum_constituicao"
}

func (v *VadeMecumConstituicao) BeforeCreate(tx *gorm.DB) error {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return nil
}

type VadeMecumConstituicaoGrupoServico struct {
	Titulo string `json:"titulo"`
}
