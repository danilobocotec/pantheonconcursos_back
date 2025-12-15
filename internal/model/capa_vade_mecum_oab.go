package model

// CapaVadeMecumOAB representa os metadados de agrupamento para o vade-mécum OAB.
type CapaVadeMecumOAB struct {
	NomeCodigo string `gorm:"column:nomecodigo;type:text;primaryKey" json:"nomecodigo"`
	Cabecalho  string `gorm:"column:Cabecalho;type:text" json:"Cabecalho"`
	Grupo      string `gorm:"column:grupo;type:text" json:"grupo"`
}

func (CapaVadeMecumOAB) TableName() string {
	return "capa_vade_mecum_oab"
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
