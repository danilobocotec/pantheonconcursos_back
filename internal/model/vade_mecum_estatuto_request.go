package model

type CreateVadeMecumEstatutoRequest struct {
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
	Artigos        string `json:"Artigos"`
	Ordem          string `json:"Ordem"`
}

type UpdateVadeMecumEstatutoRequest struct {
	IDTipo         *string `json:"idtipo"`
	Tipo           *string `json:"tipo"`
	IDCodigo       *string `json:"idcodigo"`
	NomeCodigo     *string `json:"nomecodigo"`
	Cabecalho      *string `json:"Cabecalho"`
	Parte          *string `json:"PARTE"`
	LivroID        *string `json:"idlivro"`
	Livro          *string `json:"livro"`
	LivroTexto     *string `json:"livrotexto"`
	TituloID       *string `json:"idtitulo"`
	Titulo         *string `json:"titulo"`
	TituloTexto    *string `json:"titulotexto"`
	SubtituloID    *string `json:"idsubtitulo"`
	Subtitulo      *string `json:"subtitulo"`
	SubtituloTexto *string `json:"subtitulotexto"`
	CapituloID     *string `json:"idcapitulo"`
	Capitulo       *string `json:"capitulo"`
	CapituloTexto  *string `json:"capitulotexto"`
	SecaoID        *string `json:"idsecao"`
	Secao          *string `json:"secao"`
	SecaoTexto     *string `json:"secaotexto"`
	SubsecaoID     *string `json:"idsubsecao"`
	Subsecao       *string `json:"subsecao"`
	SubsecaoTexto  *string `json:"subsecaotexto"`
	NumeroArtigo   *string `json:"num_artigo"`
	Artigos        *string `json:"Artigos"`
	Ordem          *string `json:"Ordem"`
}
