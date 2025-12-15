package model

// CapaVadeMecumJurisprudencia representa metadados de agrupamento das jurisprudências.
type CapaVadeMecumJurisprudencia struct {
	NomeCodigo string `gorm:"column:nomecodigo;type:text;primaryKey" json:"nomecodigo"`
	Cabecalho  string `gorm:"column:Cabecalho;type:text" json:"Cabecalho"`
	Grupo      string `gorm:"column:grupo;type:text" json:"grupo"`
}

func (CapaVadeMecumJurisprudencia) TableName() string {
	return "capa_vade_mecum_jurisprudencia"
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
