package model

import (
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CapaVadeMecumOAB representa os metadados de agrupamento para o vade-mécum OAB.
type CapaVadeMecumOAB struct {
	ID         string `gorm:"column:id;type:text;primaryKey" json:"id"`
	NomeCodigo string `gorm:"column:nomecodigo;type:text;unique" json:"nomecodigo"`
	Cabecalho  string `gorm:"column:Cabecalho;type:text" json:"Cabecalho"`
	Grupo      string `gorm:"column:grupo;type:text" json:"grupo"`
}

func (CapaVadeMecumOAB) TableName() string {
	return "capa_vade_mecum_oab"
}

func (c *CapaVadeMecumOAB) BeforeCreate(tx *gorm.DB) error {
	if strings.TrimSpace(c.ID) == "" {
		c.ID = uuid.NewString()
	}
	return nil
}

type CreateCapaVadeMecumOABRequest struct {
	NomeCodigo string `json:"nomecodigo" binding:"required"`
	Cabecalho  string `json:"Cabecalho"`
	Grupo      string `json:"grupo"`
}

type UpdateCapaVadeMecumOABRequest struct {
	Cabecalho *string `json:"Cabecalho"`
	Grupo     *string `json:"grupo"`
}
