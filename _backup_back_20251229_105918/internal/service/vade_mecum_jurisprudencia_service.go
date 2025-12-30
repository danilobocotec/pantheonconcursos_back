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
	"gorm.io/gorm"
)

type VadeMecumJurisprudenciaService struct {
	repo *repository.VadeMecumJurisprudenciaRepository
}

func NewVadeMecumJurisprudenciaService(repo *repository.VadeMecumJurisprudenciaRepository) *VadeMecumJurisprudenciaService {
	return &VadeMecumJurisprudenciaService{repo: repo}
}

var jurisprudenciaHeadersSplitEnunciadoLegacy = []string{
	"idtipo",
	"tipo",
	"idcodigo",
	"nomecodigo",
	"Cabecalho",
	"Tipo",
	"idramo",
	"ramotexto",
	"idassunto",
	"assuntotexto",
	"idenunciado",
	"Enunciado",
	"Enunciado1",
	"Enunciado2",
	"Enunciado3",
	"Enunciado4",
	"Enunciado5",
	"Enunciado6",
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

var jurisprudenciaHeadersSplitEnunciadoLegacyWithID = append([]string{"id"}, jurisprudenciaHeadersSplitEnunciadoLegacy...)

type jurisprudenciaImportSchema struct {
	enunciadoOffsets []int
	idSecaoOffset    int
}

var jurisprudenciaSchemaSplitEnunciadoLegacy = jurisprudenciaImportSchema{
	enunciadoOffsets: []int{11, 12, 13, 14, 15, 16, 17},
	idSecaoOffset:    18,
}

func (s *VadeMecumJurisprudenciaService) ImportFromExcel(r io.Reader) (int, error) {
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
	switch {
	case headersMatch(header, jurisprudenciaHeadersSplitEnunciadoLegacyWithID):
		return s.importRows(rows[1:], true, jurisprudenciaSchemaSplitEnunciadoLegacy)
	case headersMatch(header, jurisprudenciaHeadersSplitEnunciadoLegacy):
		return s.importRows(rows[1:], false, jurisprudenciaSchemaSplitEnunciadoLegacy)
	default:
		return 0, fmt.Errorf("cabecalho invalido: esperado %v ou %v", jurisprudenciaHeadersSplitEnunciadoLegacyWithID, jurisprudenciaHeadersSplitEnunciadoLegacy)
	}
}

func (s *VadeMecumJurisprudenciaService) importRows(rows [][]string, hasID bool, schema jurisprudenciaImportSchema) (int, error) {
	var batch []*model.VadeMecumJurisprudencia

	for idx, row := range rows {
		if isRowEmpty(row) {
			continue
		}

		offset := 0
		rawID := ""
		if hasID {
			rawID = strings.TrimSpace(getCellValue(row, 0))
			offset = 1
		}

		nomeCodigo := strings.TrimSpace(getCellValue(row, offset+3))
		if nomeCodigo == "" {
			return 0, fmt.Errorf("linha %d: nomecodigo é obrigatório", idx+2)
		}

		enunciado := buildEnunciado(row, offset, schema.enunciadoOffsets)
		idSecaoIdx := offset + schema.idSecaoOffset
		secaoIdx := idSecaoIdx + 1
		secaoTextoIdx := idSecaoIdx + 2
		idSubsecaoIdx := idSecaoIdx + 3
		subsecaoIdx := idSecaoIdx + 4
		subsecaoTextoIdx := idSecaoIdx + 5
		numArtigoIdx := idSecaoIdx + 6
		normativoIdx := idSecaoIdx + 7
		ordemIdx := idSecaoIdx + 8

		numArtigo := strings.TrimSpace(getCellValue(row, numArtigoIdx))
		normativo := strings.TrimSpace(getCellValue(row, normativoIdx))

		key := strings.ToLower(strings.Join([]string{nomeCodigo, normativo, numArtigo, enunciado}, "|"))
		deterministicID := ""
		if key != "" {
			deterministicID = uuid.NewSHA1(uuid.NameSpaceOID, []byte(key)).String()
		}

		id := rawID
		if strings.TrimSpace(id) == "" {
			if deterministicID != "" {
				id = deterministicID
			} else {
				id = uuid.NewString()
			}
		}

		item := &model.VadeMecumJurisprudencia{
			ID:            id,
			IDTipo:        getCellValue(row, offset+0),
			Tipo:          getCellValue(row, offset+1),
			IDCodigo:      getCellValue(row, offset+2),
			NomeCodigo:    nomeCodigo,
			Cabecalho:     getCellValue(row, offset+4),
			TipoDescricao: getCellValue(row, offset+5),
			IDRamo:        getCellValue(row, offset+6),
			RamoTexto:     getCellValue(row, offset+7),
			IDAssunto:     getCellValue(row, offset+8),
			AssuntoTexto:  getCellValue(row, offset+9),
			IDEnunciado:   getCellValue(row, offset+10),
			Enunciado:     enunciado,
			IDSecao:       getCellValue(row, idSecaoIdx),
			Secao:         getCellValue(row, secaoIdx),
			SecaoTexto:    getCellValue(row, secaoTextoIdx),
			IDSubsecao:    getCellValue(row, idSubsecaoIdx),
			Subsecao:      getCellValue(row, subsecaoIdx),
			SubsecaoTexto: getCellValue(row, subsecaoTextoIdx),
			NumeroArtigo:  numArtigo,
			Normativo:     normativo,
			Ordem:         getCellValue(row, ordemIdx),
		}

		batch = append(batch, item)
	}

	if len(batch) == 0 {
		return 0, errors.New("nenhuma linha válida encontrada na planilha")
	}

	unique := deduplicateJurisprudencia(batch)

	if err := s.repo.Upsert(unique); err != nil {
		return 0, fmt.Errorf("falha ao salvar registros: %w", err)
	}

	return len(unique), nil
}

func buildEnunciado(row []string, offset int, enunciadoOffsets []int) string {
	parts := make([]string, 0, len(enunciadoOffsets))
	for _, enunciadoOffset := range enunciadoOffsets {
		value := strings.TrimSpace(getCellValue(row, offset+enunciadoOffset))
		if value != "" {
			parts = append(parts, value)
		}
	}
	return strings.Join(parts, " ")
}

func (s *VadeMecumJurisprudenciaService) GetAll() ([]model.VadeMecumJurisprudencia, error) {
	return s.repo.GetAll()
}

func (s *VadeMecumJurisprudenciaService) GetGroupedByNomeCodigo() ([]model.VadeMecumJurisprudenciaGroup, error) {
	return s.repo.GetGroupedByNomeCodigo()
}

func (s *VadeMecumJurisprudenciaService) Reset() error {
	return s.repo.DeleteAll()
}

func (s *VadeMecumJurisprudenciaService) GetByID(id string) (*model.VadeMecumJurisprudencia, error) {
	trimmed := strings.TrimSpace(id)
	if trimmed == "" {
		return nil, errors.New("id inválido")
	}
	return s.repo.GetByID(trimmed)
}

func (s *VadeMecumJurisprudenciaService) Create(req *model.CreateVadeMecumJurisprudenciaRequest) (*model.VadeMecumJurisprudencia, error) {
	if req == nil {
		return nil, errors.New("payload obrigatório")
	}

	nome := strings.TrimSpace(req.NomeCodigo)
	if nome == "" {
		return nil, errors.New("nomecodigo é obrigatório")
	}

	id := strings.TrimSpace(req.ID)
	if id == "" {
		id = uuid.NewString()
	} else {
		if _, err := s.repo.GetByID(id); err == nil {
			return nil, fmt.Errorf("registro com id '%s' já existe", id)
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	item := &model.VadeMecumJurisprudencia{
		ID:            id,
		IDTipo:        strings.TrimSpace(req.IDTipo),
		Tipo:          strings.TrimSpace(req.Tipo),
		IDCodigo:      strings.TrimSpace(req.IDCodigo),
		NomeCodigo:    nome,
		Cabecalho:     strings.TrimSpace(req.Cabecalho),
		TipoDescricao: strings.TrimSpace(req.TipoDescricao),
		IDRamo:        strings.TrimSpace(req.IDRamo),
		RamoTexto:     strings.TrimSpace(req.RamoTexto),
		IDAssunto:     strings.TrimSpace(req.IDAssunto),
		AssuntoTexto:  strings.TrimSpace(req.AssuntoTexto),
		IDEnunciado:   strings.TrimSpace(req.IDEnunciado),
		Enunciado:     strings.TrimSpace(req.Enunciado),
		IDSecao:       strings.TrimSpace(req.IDSecao),
		Secao:         strings.TrimSpace(req.Secao),
		SecaoTexto:    strings.TrimSpace(req.SecaoTexto),
		IDSubsecao:    strings.TrimSpace(req.IDSubsecao),
		Subsecao:      strings.TrimSpace(req.Subsecao),
		SubsecaoTexto: strings.TrimSpace(req.SubsecaoTexto),
		NumeroArtigo:  strings.TrimSpace(req.NumeroArtigo),
		Normativo:     strings.TrimSpace(req.Normativo),
		Ordem:         strings.TrimSpace(req.Ordem),
	}

	if err := s.repo.Create(item); err != nil {
		return nil, err
	}

	return item, nil
}

func (s *VadeMecumJurisprudenciaService) Update(id string, req *model.UpdateVadeMecumJurisprudenciaRequest) (*model.VadeMecumJurisprudencia, error) {
	if req == nil {
		return nil, errors.New("payload obrigatório")
	}

	trimmed := strings.TrimSpace(id)
	if trimmed == "" {
		return nil, errors.New("id inválido")
	}

	existing, err := s.repo.GetByID(trimmed)
	if err != nil {
		return nil, err
	}

	if req.IDTipo != nil {
		existing.IDTipo = strings.TrimSpace(*req.IDTipo)
	}
	if req.Tipo != nil {
		existing.Tipo = strings.TrimSpace(*req.Tipo)
	}
	if req.IDCodigo != nil {
		existing.IDCodigo = strings.TrimSpace(*req.IDCodigo)
	}
	if req.NomeCodigo != nil {
		nome := strings.TrimSpace(*req.NomeCodigo)
		if nome == "" {
			return nil, errors.New("nomecodigo não pode ser vazio")
		}
		existing.NomeCodigo = nome
	}
	if req.Cabecalho != nil {
		existing.Cabecalho = strings.TrimSpace(*req.Cabecalho)
	}
	if req.TipoDescricao != nil {
		existing.TipoDescricao = strings.TrimSpace(*req.TipoDescricao)
	}
	if req.IDRamo != nil {
		existing.IDRamo = strings.TrimSpace(*req.IDRamo)
	}
	if req.RamoTexto != nil {
		existing.RamoTexto = strings.TrimSpace(*req.RamoTexto)
	}
	if req.IDAssunto != nil {
		existing.IDAssunto = strings.TrimSpace(*req.IDAssunto)
	}
	if req.AssuntoTexto != nil {
		existing.AssuntoTexto = strings.TrimSpace(*req.AssuntoTexto)
	}
	if req.IDEnunciado != nil {
		existing.IDEnunciado = strings.TrimSpace(*req.IDEnunciado)
	}
	if req.Enunciado != nil {
		existing.Enunciado = strings.TrimSpace(*req.Enunciado)
	}
	if req.IDSecao != nil {
		existing.IDSecao = strings.TrimSpace(*req.IDSecao)
	}
	if req.Secao != nil {
		existing.Secao = strings.TrimSpace(*req.Secao)
	}
	if req.SecaoTexto != nil {
		existing.SecaoTexto = strings.TrimSpace(*req.SecaoTexto)
	}
	if req.IDSubsecao != nil {
		existing.IDSubsecao = strings.TrimSpace(*req.IDSubsecao)
	}
	if req.Subsecao != nil {
		existing.Subsecao = strings.TrimSpace(*req.Subsecao)
	}
	if req.SubsecaoTexto != nil {
		existing.SubsecaoTexto = strings.TrimSpace(*req.SubsecaoTexto)
	}
	if req.NumeroArtigo != nil {
		existing.NumeroArtigo = strings.TrimSpace(*req.NumeroArtigo)
	}
	if req.Normativo != nil {
		existing.Normativo = strings.TrimSpace(*req.Normativo)
	}
	if req.Ordem != nil {
		existing.Ordem = strings.TrimSpace(*req.Ordem)
	}

	if err := s.repo.Update(existing); err != nil {
		return nil, err
	}

	return existing, nil
}

func (s *VadeMecumJurisprudenciaService) Delete(id string) error {
	trimmed := strings.TrimSpace(id)
	if trimmed == "" {
		return errors.New("id inválido")
	}
	return s.repo.Delete(trimmed)
}

func deduplicateJurisprudencia(items []*model.VadeMecumJurisprudencia) []*model.VadeMecumJurisprudencia {
	index := make(map[string]int, len(items))
	unique := make([]*model.VadeMecumJurisprudencia, 0, len(items))

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


















