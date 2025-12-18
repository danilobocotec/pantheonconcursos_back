package model

type CreateVadeMecumConstituicaoRequest struct {
	RegistroID    string `json:"registro_id"`
	IDTipo        string `json:"idtipo"`
	Tipo          string `json:"tipo"`
	Cabecalho     string `json:"cabecalho"`
	IDTitulo      string `json:"idtitulo"`
	Titulo        string `json:"titulo"`
	TextoDoTitulo string `json:"textodotitulo"`
	IDCapitulo    string `json:"idcapitulo"`
	Capitulo      string `json:"capitulo"`
	TextoCapitulo string `json:"textocapitulo"`
	IDSecao       string `json:"idsecao"`
	Secao         string `json:"secao"`
	TextoSecao    string `json:"textosecao"`
	IDSubsecao    string `json:"idsubsecao"`
	Subsecao      string `json:"subsecao"`
	TextoSubsecao string `json:"subsecaotexto"`
	Normativo     string `json:"Normativo"`
}

type UpdateVadeMecumConstituicaoRequest struct {
	RegistroID    *string `json:"registro_id"`
	IDTipo        *string `json:"idtipo"`
	Tipo          *string `json:"tipo"`
	Cabecalho     *string `json:"cabecalho"`
	IDTitulo      *string `json:"idtitulo"`
	Titulo        *string `json:"titulo"`
	TextoDoTitulo *string `json:"textodotitulo"`
	IDCapitulo    *string `json:"idcapitulo"`
	Capitulo      *string `json:"capitulo"`
	TextoCapitulo *string `json:"textocapitulo"`
	IDSecao       *string `json:"idsecao"`
	Secao         *string `json:"secao"`
	TextoSecao    *string `json:"textosecao"`
	IDSubsecao    *string `json:"idsubsecao"`
	Subsecao      *string `json:"subsecao"`
	TextoSubsecao *string `json:"subsecaotexto"`
	Normativo     *string `json:"Normativo"`
}
