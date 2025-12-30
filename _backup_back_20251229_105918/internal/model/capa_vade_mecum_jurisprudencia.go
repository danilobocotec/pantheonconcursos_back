package model

import (
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CapaVadeMecumJurisprudencia representa metadados de agrupamento das jurisprudências.
type CapaVadeMecumJurisprudencia struct {
	ID         string `gorm:"column:id;type:text;primaryKey" json:"id"`
	NomeCodigo string `gorm:"column:nomecodigo;type:text;unique" json:"nomecodigo"`
	Cabecalho  string `gorm:"column:Cabecalho;type:text" json:"Cabecalho"`
	Grupo      string `gorm:"column:grupo;type:text" json:"grupo"`
}

func (CapaVadeMecumJurisprudencia) TableName() string {
	return "capa_vade_mecum_jurisprudencia"
}

func (c *CapaVadeMecumJurisprudencia) BeforeCreate(tx *gorm.DB) error {
	if strings.TrimSpace(c.ID) == "" {
		c.ID = uuid.NewString()
	}
	return nil
}

type CreateCapaVadeMecumJurisprudenciaRequest struct {
	NomeCodigo string `json:"nomecodigo" binding:"required"`
	Cabecalho  string `json:"Cabecalho"`
	Grupo      string `json:"grupo"`
}

type UpdateCapaVadeMecumJurisprudenciaRequest struct {
	Cabecalho *string `json:"Cabecalho"`
	Grupo     *string `json:"grupo"`
}
