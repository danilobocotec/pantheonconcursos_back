package model

import (
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CapaVadeMecumCodigo represents the cover grouping metadata for vade-mécum códigos.
type CapaVadeMecumCodigo struct {
	ID         string `gorm:"column:id;type:text;primaryKey" json:"id"`
	NomeCodigo string `gorm:"column:nomecodigo;type:text;unique" json:"nomecodigo"`
	Cabecalho  string `gorm:"column:Cabecalho;type:text" json:"Cabecalho"`
	Grupo      string `gorm:"column:grupo;type:text" json:"grupo"`
}

func (CapaVadeMecumCodigo) TableName() string {
	return "capa_vade_mecum_codigo"
}

func (c *CapaVadeMecumCodigo) BeforeCreate(tx *gorm.DB) error {
	if strings.TrimSpace(c.ID) == "" {
		c.ID = uuid.NewString()
	}
	return nil
}

type CreateCapaVadeMecumCodigoRequest struct {
	NomeCodigo string `json:"nomecodigo" binding:"required"`
	Cabecalho  string `json:"Cabecalho"`
	Grupo      string `json:"grupo"`
}

type UpdateCapaVadeMecumCodigoRequest struct {
	Grupo *string `json:"grupo"`
}
