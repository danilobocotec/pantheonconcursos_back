package model

// CapaVadeMecumCodigo represents the cover grouping metadata for vade-mécum códigos.
type CapaVadeMecumCodigo struct {
	NomeCodigo string `gorm:"column:nomecodigo;type:text;primaryKey" json:"nomecodigo"`
	Cabecalho  string `gorm:"column:Cabecalho;type:text" json:"Cabecalho"`
	Grupo      string `gorm:"column:grupo;type:text" json:"grupo"`
}

func (CapaVadeMecumCodigo) TableName() string {
	return "capa_vade_mecum_codigo"
}

type CreateCapaVadeMecumCodigoRequest struct {
	NomeCodigo string `json:"nomecodigo" binding:"required"`
	Cabecalho  string `json:"Cabecalho"`
	Grupo      string `json:"grupo"`
}

type UpdateCapaVadeMecumCodigoRequest struct {
	Cabecalho *string `json:"Cabecalho"`
	Grupo     *string `json:"grupo"`
}
