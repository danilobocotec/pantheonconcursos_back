package service

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/google/uuid"
	"github.com/thepantheon/api/internal/model"
	"github.com/thepantheon/api/internal/repository"
	"github.com/xuri/excelize/v2"
)

type VadeMecumEstatutoService struct {
	repo *repository.VadeMecumEstatutoRepository
}

func NewVadeMecumEstatutoService(repo *repository.VadeMecumEstatutoRepository) *VadeMecumEstatutoService {
	return &VadeMecumEstatutoService{repo: repo}
}

const estatutoArtigosLimit = 5000

var estatutoImportHeaders = []string{
	"idtipo",
	"tipo",
	"idcodigo",
	"nomecodigo",
	"Cabecalho",
	"PARTE",
	"idlivro",
	"livro",
	"livrotexto",
	"idtitulo",
	"titulo",
	"titulotexto",
	"idsubtitulo",
	"subtitulo",
	"subtitulotexto",
	"idcapitulo",
	"capitulo",
	"capitulotexto",
	"idsecao",
	"secao",
	"secaotexto",
	"idsubsecao",
	"subsecao",
	"subsecaotexto",
	"num_artigo",
	"Artigos",
	"Ordem",
}

func (s *VadeMecumEstatutoService) ImportEstatuto(r io.Reader) (int, error) {
	f, err := excelize.OpenReader(r)
	if err != nil {
		return 0, fmt.Errorf("falha ao abrir planilha: %w", err)
	}
	defer f.Close()

	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return 0, errors.New("planilha sem abas")
	}

	rows, err := f.GetRows(sheets[0])
	if err != nil {
		return 0, fmt.Errorf("falha ao ler linhas: %w", err)
	}

	if len(rows) <= 1 {
		return 0, errors.New("planilha n\xE3o possui dados al\xE9m do cabe\xE7alho")
	}

	header := normalizeHeader(rows[0])
	if !headersMatch(header, estatutoImportHeaders) {
		return 0, fmt.Errorf("cabe\xE7alho inv\xE1lido: esperado %v", estatutoImportHeaders)
	}

	var batch []*model.VadeMecumEstatuto

	for _, row := range rows[1:] {
		if isRowEmpty(row) {
			continue
		}

		idCodigo := strings.TrimSpace(getCellValue(row, 2))
		if idCodigo == "" {
			idCodigo = uuid.NewString()
		}

		item := &model.VadeMecumEstatuto{
			IDTipo:         getCellValue(row, 0),
			Tipo:           getCellValue(row, 1),
			IDCodigo:       idCodigo,
			NomeCodigo:     getCellValue(row, 3),
			Cabecalho:      getCellValue(row, 4),
			Parte:          getCellValue(row, 5),
			LivroID:        getCellValue(row, 6),
			Livro:          getCellValue(row, 7),
			LivroTexto:     getCellValue(row, 8),
			TituloID:       getCellValue(row, 9),
			Titulo:         getCellValue(row, 10),
			TituloTexto:    getCellValue(row, 11),
			SubtituloID:    getCellValue(row, 12),
			Subtitulo:      getCellValue(row, 13),
			SubtituloTexto: getCellValue(row, 14),
			CapituloID:     getCellValue(row, 15),
			Capitulo:       getCellValue(row, 16),
			CapituloTexto:  getCellValue(row, 17),
			SecaoID:        getCellValue(row, 18),
			Secao:          getCellValue(row, 19),
			SecaoTexto:     getCellValue(row, 20),
			SubsecaoID:     getCellValue(row, 21),
			Subsecao:       getCellValue(row, 22),
			SubsecaoTexto:  getCellValue(row, 23),
			NumeroArtigo:   getCellValue(row, 24),
			Artigos:        getCellValue(row, 25),
			Ordem:          getCellValue(row, 26),
		}

		batch = append(batch, item)
	}

	if len(batch) == 0 {
		return 0, errors.New("nenhuma linha v\xE1lida encontrada na planilha")
	}

	if err := s.repo.UpsertByCodigo(batch); err != nil {
		return 0, fmt.Errorf("falha ao salvar registros: %w", err)
	}

	return len(batch), nil
}

func (s *VadeMecumEstatutoService) GetAll() ([]model.VadeMecumEstatuto, error) {
	return s.repo.GetAll()
}

func (s *VadeMecumEstatutoService) GrupoServico() ([]model.VadeMecumEstatutoGrupoServico, error) {
	return s.repo.GetGrupoServico()
}

func (s *VadeMecumEstatutoService) GetByID(id uuid.UUID) (*model.VadeMecumEstatuto, error) {
	return s.repo.GetByID(id)
}

func (s *VadeMecumEstatutoService) Create(req *model.CreateVadeMecumEstatutoRequest) (*model.VadeMecumEstatuto, error) {
	if req == nil {
		return nil, errors.New("payload obrigatório")
	}

	idCodigo := strings.TrimSpace(req.IDCodigo)
	if idCodigo == "" {
		idCodigo = uuid.NewString()
	}

	item := &model.VadeMecumEstatuto{
		IDTipo:         strings.TrimSpace(req.IDTipo),
		Tipo:           strings.TrimSpace(req.Tipo),
		IDCodigo:       idCodigo,
		NomeCodigo:     strings.TrimSpace(req.NomeCodigo),
		Cabecalho:      strings.TrimSpace(req.Cabecalho),
		Parte:          strings.TrimSpace(req.Parte),
		LivroID:        strings.TrimSpace(req.LivroID),
		Livro:          strings.TrimSpace(req.Livro),
		LivroTexto:     strings.TrimSpace(req.LivroTexto),
		TituloID:       strings.TrimSpace(req.TituloID),
		Titulo:         strings.TrimSpace(req.Titulo),
		TituloTexto:    strings.TrimSpace(req.TituloTexto),
		SubtituloID:    strings.TrimSpace(req.SubtituloID),
		Subtitulo:      strings.TrimSpace(req.Subtitulo),
		SubtituloTexto: strings.TrimSpace(req.SubtituloTexto),
		CapituloID:     strings.TrimSpace(req.CapituloID),
		Capitulo:       strings.TrimSpace(req.Capitulo),
		CapituloTexto:  strings.TrimSpace(req.CapituloTexto),
		SecaoID:        strings.TrimSpace(req.SecaoID),
		Secao:          strings.TrimSpace(req.Secao),
		SecaoTexto:     strings.TrimSpace(req.SecaoTexto),
		SubsecaoID:     strings.TrimSpace(req.SubsecaoID),
		Subsecao:       strings.TrimSpace(req.Subsecao),
		SubsecaoTexto:  strings.TrimSpace(req.SubsecaoTexto),
		NumeroArtigo:   strings.TrimSpace(req.NumeroArtigo),
		Artigos:        strings.TrimSpace(req.Artigos),
		Ordem:          strings.TrimSpace(req.Ordem),
	}

	if err := s.repo.Create(item); err != nil {
		return nil, err
	}

	return item, nil
}

func (s *VadeMecumEstatutoService) Update(id uuid.UUID, req *model.UpdateVadeMecumEstatutoRequest) (*model.VadeMecumEstatuto, error) {
	if req == nil {
		return nil, errors.New("payload obrigatório")
	}

	item, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.IDTipo != nil {
		item.IDTipo = strings.TrimSpace(*req.IDTipo)
	}
	if req.Tipo != nil {
		item.Tipo = strings.TrimSpace(*req.Tipo)
	}
	if req.IDCodigo != nil {
		item.IDCodigo = strings.TrimSpace(*req.IDCodigo)
	}
	if req.NomeCodigo != nil {
		item.NomeCodigo = strings.TrimSpace(*req.NomeCodigo)
	}
	if req.Cabecalho != nil {
		item.Cabecalho = strings.TrimSpace(*req.Cabecalho)
	}
	if req.Parte != nil {
		item.Parte = strings.TrimSpace(*req.Parte)
	}
	if req.LivroID != nil {
		item.LivroID = strings.TrimSpace(*req.LivroID)
	}
	if req.Livro != nil {
		item.Livro = strings.TrimSpace(*req.Livro)
	}
	if req.LivroTexto != nil {
		item.LivroTexto = strings.TrimSpace(*req.LivroTexto)
	}
	if req.TituloID != nil {
		item.TituloID = strings.TrimSpace(*req.TituloID)
	}
	if req.Titulo != nil {
		item.Titulo = strings.TrimSpace(*req.Titulo)
	}
	if req.TituloTexto != nil {
		item.TituloTexto = strings.TrimSpace(*req.TituloTexto)
	}
	if req.SubtituloID != nil {
		item.SubtituloID = strings.TrimSpace(*req.SubtituloID)
	}
	if req.Subtitulo != nil {
		item.Subtitulo = strings.TrimSpace(*req.Subtitulo)
	}
	if req.SubtituloTexto != nil {
		item.SubtituloTexto = strings.TrimSpace(*req.SubtituloTexto)
	}
	if req.CapituloID != nil {
		item.CapituloID = strings.TrimSpace(*req.CapituloID)
	}
	if req.Capitulo != nil {
		item.Capitulo = strings.TrimSpace(*req.Capitulo)
	}
	if req.CapituloTexto != nil {
		item.CapituloTexto = strings.TrimSpace(*req.CapituloTexto)
	}
	if req.SecaoID != nil {
		item.SecaoID = strings.TrimSpace(*req.SecaoID)
	}
	if req.Secao != nil {
		item.Secao = strings.TrimSpace(*req.Secao)
	}
	if req.SecaoTexto != nil {
		item.SecaoTexto = strings.TrimSpace(*req.SecaoTexto)
	}
	if req.SubsecaoID != nil {
		item.SubsecaoID = strings.TrimSpace(*req.SubsecaoID)
	}
	if req.Subsecao != nil {
		item.Subsecao = strings.TrimSpace(*req.Subsecao)
	}
	if req.SubsecaoTexto != nil {
		item.SubsecaoTexto = strings.TrimSpace(*req.SubsecaoTexto)
	}
	if req.NumeroArtigo != nil {
		item.NumeroArtigo = strings.TrimSpace(*req.NumeroArtigo)
	}
	if req.Artigos != nil {
		item.Artigos = strings.TrimSpace(*req.Artigos)
	}
	if req.Ordem != nil {
		item.Ordem = strings.TrimSpace(*req.Ordem)
	}

	if err := s.repo.Update(item); err != nil {
		return nil, err
	}

	return item, nil
}

func (s *VadeMecumEstatutoService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
