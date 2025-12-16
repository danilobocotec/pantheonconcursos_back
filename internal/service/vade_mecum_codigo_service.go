package service

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/google/uuid"
	"github.com/thepantheon/api/internal/model"
	"github.com/thepantheon/api/internal/repository"
	"github.com/xuri/excelize/v2"
)

type VadeMecumCodigoService struct {
	repo *repository.VadeMecumCodigoRepository
}

func NewVadeMecumCodigoService(repo *repository.VadeMecumCodigoRepository) *VadeMecumCodigoService {
	return &VadeMecumCodigoService{repo: repo}
}

var vadeMecumImportHeaders = []string{
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
	"Normativo",
	"Ordem",
}

<<<<<<< HEAD

func (s *VadeMecumCodigoService) Create(req *model.CreateVadeMecumCodigoRequest) (*model.VadeMecumCodigo, error) {
	idCodigo := strings.TrimSpace(req.IDCodigo)
	if idCodigo == "" {
		idCodigo = uuid.NewString()
	}

	item := &model.VadeMecumCodigo{
		IDTipo:         req.IDTipo,
		Tipo:           req.Tipo,
		IDCodigo:       idCodigo,
=======
func (s *VadeMecumCodigoService) Create(req *model.CreateVadeMecumCodigoRequest) (*model.VadeMecumCodigo, error) {
	item := &model.VadeMecumCodigo{
		IDTipo:         req.IDTipo,
		Tipo:           req.Tipo,
		IDCodigo:       req.IDCodigo,
>>>>>>> 451427c4618a62b6f9ac9376f15b00d127a565e5
		NomeCodigo:     req.NomeCodigo,
		Cabecalho:      req.Cabecalho,
		Parte:          req.Parte,
		LivroID:        req.LivroID,
		Livro:          req.Livro,
		LivroTexto:     req.LivroTexto,
		TituloID:       req.TituloID,
		Titulo:         req.Titulo,
		TituloTexto:    req.TituloTexto,
		SubtituloID:    req.SubtituloID,
		Subtitulo:      req.Subtitulo,
		SubtituloTexto: req.SubtituloTexto,
		CapituloID:     req.CapituloID,
		Capitulo:       req.Capitulo,
		CapituloTexto:  req.CapituloTexto,
		SecaoID:        req.SecaoID,
		Secao:          req.Secao,
		SecaoTexto:     req.SecaoTexto,
		SubsecaoID:     req.SubsecaoID,
		Subsecao:       req.Subsecao,
		SubsecaoTexto:  req.SubsecaoTexto,
		NumeroArtigo:   req.NumeroArtigo,
		Normativo:      req.Normativo,
		Ordem:          req.Ordem,
	}

	if err := s.repo.Create(item); err != nil {
		return nil, err
	}
	return item, nil
}

func (s *VadeMecumCodigoService) GetAll() ([]model.VadeMecumCodigo, error) {
	return s.repo.GetAll()
}

<<<<<<< HEAD

=======
>>>>>>> 451427c4618a62b6f9ac9376f15b00d127a565e5
func (s *VadeMecumCodigoService) GetByID(id uuid.UUID) (*model.VadeMecumCodigo, error) {
	return s.repo.GetByID(id)
}

func (s *VadeMecumCodigoService) Update(id uuid.UUID, req *model.UpdateVadeMecumCodigoRequest) (*model.VadeMecumCodigo, error) {
	item, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.IDTipo != "" {
		item.IDTipo = req.IDTipo
	}
	if req.Tipo != "" {
		item.Tipo = req.Tipo
	}
	if req.IDCodigo != "" {
		item.IDCodigo = req.IDCodigo
	}
	if req.NomeCodigo != "" {
		item.NomeCodigo = req.NomeCodigo
	}
	if req.Cabecalho != "" {
		item.Cabecalho = req.Cabecalho
	}
	if req.Parte != "" {
		item.Parte = req.Parte
	}
	if req.LivroID != "" {
		item.LivroID = req.LivroID
	}
	if req.Livro != "" {
		item.Livro = req.Livro
	}
	if req.LivroTexto != "" {
		item.LivroTexto = req.LivroTexto
	}
	if req.TituloID != "" {
		item.TituloID = req.TituloID
	}
	if req.Titulo != "" {
		item.Titulo = req.Titulo
	}
	if req.TituloTexto != "" {
		item.TituloTexto = req.TituloTexto
	}
	if req.SubtituloID != "" {
		item.SubtituloID = req.SubtituloID
	}
	if req.Subtitulo != "" {
		item.Subtitulo = req.Subtitulo
	}
	if req.SubtituloTexto != "" {
		item.SubtituloTexto = req.SubtituloTexto
	}
	if req.CapituloID != "" {
		item.CapituloID = req.CapituloID
	}
	if req.Capitulo != "" {
		item.Capitulo = req.Capitulo
	}
	if req.CapituloTexto != "" {
		item.CapituloTexto = req.CapituloTexto
	}
	if req.SecaoID != "" {
		item.SecaoID = req.SecaoID
	}
	if req.Secao != "" {
		item.Secao = req.Secao
	}
	if req.SecaoTexto != "" {
		item.SecaoTexto = req.SecaoTexto
	}
	if req.SubsecaoID != "" {
		item.SubsecaoID = req.SubsecaoID
	}
	if req.Subsecao != "" {
		item.Subsecao = req.Subsecao
	}
	if req.SubsecaoTexto != "" {
		item.SubsecaoTexto = req.SubsecaoTexto
	}
	if req.NumeroArtigo != "" {
		item.NumeroArtigo = req.NumeroArtigo
	}
	if req.Normativo != "" {
		item.Normativo = req.Normativo
	}
	if req.Ordem != "" {
		item.Ordem = req.Ordem
	}

	if err := s.repo.Update(item); err != nil {
		return nil, err
	}

	return item, nil
}

func (s *VadeMecumCodigoService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *VadeMecumCodigoService) ImportFromExcel(r io.Reader) (int, error) {
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
		return 0, errors.New("planilha não possui dados além do cabeçalho")
	}

<<<<<<< HEAD
	header := normalizeHeader(rows[0])

	switch {
	case headersMatch(header, vadeMecumImportHeaders):
		return s.importLegacyRows(rows[1:])
	case headersMatch(header, vadeMecumImportHeadersV2):
		return s.importV2Rows(rows[1:])
	default:
		return 0, fmt.Errorf("cabeçalho inválido: esperado formatos %v ou %v", vadeMecumImportHeaders, vadeMecumImportHeadersV2)
	}
}

func (s *VadeMecumCodigoService) importLegacyRows(rows [][]string) (int, error) {
	var batch []*model.VadeMecumCodigo

	for idx, row := range rows {
=======
	if err := validateImportHeader(rows[0]); err != nil {
		return 0, err
	}

	var batch []*model.VadeMecumCodigo
	for idx, row := range rows[1:] {
>>>>>>> 451427c4618a62b6f9ac9376f15b00d127a565e5
		if isRowEmpty(row) {
			continue
		}

<<<<<<< HEAD
		idCodigo := strings.TrimSpace(getCellValue(row, 2))
		if idCodigo == "" {
			idCodigo = uuid.NewString()
		}

		item := &model.VadeMecumCodigo{
			IDTipo:         getCellValue(row, 0),
			Tipo:           getCellValue(row, 1),
			IDCodigo:       idCodigo,
=======
		item := &model.VadeMecumCodigo{
			IDTipo:         getCellValue(row, 0),
			Tipo:           getCellValue(row, 1),
			IDCodigo:       getCellValue(row, 2),
>>>>>>> 451427c4618a62b6f9ac9376f15b00d127a565e5
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
			Normativo:      getCellValue(row, 25),
			Ordem:          getCellValue(row, 26),
		}

		if strings.TrimSpace(item.NomeCodigo) == "" {
			return 0, fmt.Errorf("linha %d: nomecodigo é obrigatório", idx+2)
		}

		batch = append(batch, item)
	}

	if len(batch) == 0 {
		return 0, errors.New("nenhuma linha válida encontrada na planilha")
	}

<<<<<<< HEAD
	uniqueBatch := deduplicateCodigos(batch)

	if err := s.repo.UpsertByCodigo(uniqueBatch); err != nil {
		return 0, fmt.Errorf("falha ao salvar registros: %w", err)
	}

	return len(uniqueBatch), nil
}

func (s *VadeMecumCodigoService) importV2Rows(rows [][]string) (int, error) {
	var batch []*model.VadeMecumCodigo

	for idx, row := range rows {
		if isRowEmpty(row) {
			continue
		}

		idCodigo := strings.TrimSpace(getCellValue(row, 0))
		if idCodigo == "" {
			idCodigo = uuid.NewString()
		}

		item := &model.VadeMecumCodigo{
			IDCodigo:      idCodigo,
			IDTipo:        getCellValue(row, 1),
			Tipo:          getCellValue(row, 2),
			NomeCodigo:    getCellValue(row, 3),
			Cabecalho:     getCellValue(row, 4),
			LivroID:       getCellValue(row, 5),
			Livro:         getCellValue(row, 6),
			LivroTexto:    getCellValue(row, 7),
			TituloID:      getCellValue(row, 8),
			Titulo:        getCellValue(row, 9),
			TituloTexto:   getCellValue(row, 10),
			CapituloID:    getCellValue(row, 11),
			Capitulo:      getCellValue(row, 12),
			CapituloTexto: getCellValue(row, 13),
			SecaoID:       getCellValue(row, 14),
			Secao:         getCellValue(row, 15),
			SecaoTexto:    getCellValue(row, 16),
			SubsecaoID:    getCellValue(row, 17),
			Subsecao:      getCellValue(row, 18),
			SubsecaoTexto: getCellValue(row, 19),
			NumeroArtigo:  getCellValue(row, 20),
			Normativo:     getCellValue(row, 21),
			Ordem:         getCellValue(row, 22),
		}

		if strings.TrimSpace(item.NomeCodigo) == "" {
			return 0, fmt.Errorf("linha %d: nomecodigo é obrigatório", idx+2)
		}

		batch = append(batch, item)
	}

	if len(batch) == 0 {
		return 0, errors.New("nenhuma linha válida encontrada na planilha")
	}

	uniqueBatch := deduplicateCodigos(batch)

	if err := s.repo.UpsertByCodigo(uniqueBatch); err != nil {
		return 0, fmt.Errorf("falha ao salvar registros: %w", err)
	}

	return len(uniqueBatch), nil
}

func normalizeHeader(row []string) []string {
	normalized := make([]string, len(row))
	for i, value := range row {
		normalized[i] = strings.TrimSpace(value)
	}
	return normalized
}

func headersMatch(header []string, expected []string) bool {
	if len(header) < len(expected) {
		return false
	}

	for idx, exp := range expected {
		if header[idx] != exp {
			return false
		}
	}

	// Ensure there are no unexpected extra columns with content where we already mapped all expected
	for _, value := range header[len(expected):] {
		if value != "" {
			return false
		}
	}

	return true
=======
	if err := s.repo.UpsertByCodigo(batch); err != nil {
		return 0, fmt.Errorf("falha ao salvar registros: %w", err)
	}

	return len(batch), nil
}

func validateImportHeader(header []string) error {
	if len(header) < len(vadeMecumImportHeaders) {
		return fmt.Errorf("cabeçalho inválido: esperado %d colunas, recebido %d", len(vadeMecumImportHeaders), len(header))
	}
	for idx, expected := range vadeMecumImportHeaders {
		cell := strings.TrimSpace(header[idx])
		if cell != expected {
			return fmt.Errorf("cabeçalho inválido na coluna %d: esperado '%s', recebido '%s'", idx+1, expected, cell)
		}
	}
	return nil
>>>>>>> 451427c4618a62b6f9ac9376f15b00d127a565e5
}

func getCellValue(row []string, idx int) string {
	if idx < len(row) {
		return strings.TrimSpace(row[idx])
	}
	return ""
}

func isRowEmpty(row []string) bool {
	for _, v := range row {
		if strings.TrimSpace(v) != "" {
			return false
		}
	}
	return true
}
<<<<<<< HEAD

func deduplicateCodigos(items []*model.VadeMecumCodigo) []*model.VadeMecumCodigo {
	index := make(map[string]int, len(items))
	unique := make([]*model.VadeMecumCodigo, 0, len(items))

	for _, item := range items {
		key := strings.TrimSpace(item.IDCodigo)
		if key == "" {
			unique = append(unique, item)
			continue
		}

		if pos, ok := index[key]; ok {
			unique[pos] = item
		} else {
			index[key] = len(unique)
			unique = append(unique, item)
		}
	}

	return unique
}

func intPtr(value int) *int {
	return &value
}
=======
>>>>>>> 451427c4618a62b6f9ac9376f15b00d127a565e5
