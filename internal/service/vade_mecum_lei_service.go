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

type VadeMecumLeiService struct {
	repo *repository.VadeMecumLeiRepository
}

func NewVadeMecumLeiService(repo *repository.VadeMecumLeiRepository) *VadeMecumLeiService {
	return &VadeMecumLeiService{repo: repo}
}

var vadeMecumLeisHeaders = []string{
	"id",
	"idtipo",
	"tipo",
	"nomecodigo",
	"Cabecalho",
	"idPARTE",
	"PARTE",
	"PARTETEXTO",
	"idtitulo",
	"titulo",
	"titulotexto",
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

func (s *VadeMecumLeiService) ImportFromExcel(r io.Reader) (int, error) {
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

	header := normalizeHeader(rows[0])
	if !headersMatch(header, vadeMecumLeisHeaders) {
		return 0, fmt.Errorf("cabeçalho inválido: esperado %v", vadeMecumLeisHeaders)
	}

	var batch []*model.VadeMecumLei
	for idx, row := range rows[1:] {
		if isRowEmpty(row) {
			continue
		}

		id := strings.TrimSpace(getCellValue(row, 0))
		if id == "" {
			id = uuid.NewString()
		}

		item := &model.VadeMecumLei{
			ID:            id,
			IDTipo:        getCellValue(row, 1),
			Tipo:          getCellValue(row, 2),
			NomeCodigo:    getCellValue(row, 3),
			Cabecalho:     getCellValue(row, 4),
			IDParte:       getCellValue(row, 5),
			Parte:         getCellValue(row, 6),
			ParteTexto:    getCellValue(row, 7),
			IDTitulo:      getCellValue(row, 8),
			Titulo:        getCellValue(row, 9),
			TituloTexto:   getCellValue(row, 10),
			IDCapitulo:    getCellValue(row, 11),
			Capitulo:      getCellValue(row, 12),
			CapituloTexto: getCellValue(row, 13),
			IDSecao:       getCellValue(row, 14),
			Secao:         getCellValue(row, 15),
			SecaoTexto:    getCellValue(row, 16),
			IDSubsecao:    getCellValue(row, 17),
			Subsecao:      getCellValue(row, 18),
			SubsecaoTexto: getCellValue(row, 19),
			NumeroArtigo:  getCellValue(row, 20),
			Artigos:       getCellValue(row, 21),
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

	unique := deduplicateLeis(batch)

	if err := s.repo.Upsert(unique); err != nil {
		return 0, fmt.Errorf("falha ao salvar registros: %w", err)
	}

	return len(unique), nil
}

func deduplicateLeis(items []*model.VadeMecumLei) []*model.VadeMecumLei {
	index := make(map[string]int, len(items))
	unique := make([]*model.VadeMecumLei, 0, len(items))

	for _, item := range items {
		key := strings.TrimSpace(item.ID)
		if key == "" {
			key = uuid.NewString()
			item.ID = key
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
