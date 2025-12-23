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

type VadeMecumOABService struct {
	repo *repository.VadeMecumOABRepository
}

func NewVadeMecumOABService(repo *repository.VadeMecumOABRepository) *VadeMecumOABService {
	return &VadeMecumOABService{repo: repo}
}

var vadeMecumOABHeaders = []string{
	"idtipo",
	"tipo",
	"nomecodigo",
	"Cabecalho",
	"titulo",
	"titulotexto",
	"TÍTULO",
	"capitulo",
	"capitulotexto",
	"CAPÍTULO",
	"secao",
	"secaotexto",
	"Seção",
	"subsecao",
	"subsecaotexto",
	"Subseção",
	"num_artigo",
	"Artigos",
}

var vadeMecumOABHeadersWithID = append([]string{"id"}, vadeMecumOABHeaders...)

func (s *VadeMecumOABService) ImportFromExcel(r io.Reader) (int, error) {
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
	case headersMatch(header, vadeMecumOABHeadersWithID):
		return s.importOABRows(rows[1:], true)
	case headersMatch(header, vadeMecumOABHeaders):
		return s.importOABRows(rows[1:], false)
	default:
		return 0, fmt.Errorf("cabeçalho inválido: esperado %v ou %v", vadeMecumOABHeadersWithID, vadeMecumOABHeaders)
	}
}

func (s *VadeMecumOABService) importOABRows(rows [][]string, hasID bool) (int, error) {
	var batch []*model.VadeMecumOAB

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

		nomeCodigo := strings.TrimSpace(getCellValue(row, offset+2))
		if nomeCodigo == "" {
			return 0, fmt.Errorf("linha %d: nomecodigo é obrigatório", idx+2)
		}

		numeroArtigo := strings.TrimSpace(getCellValue(row, offset+16))
		artigos := strings.TrimSpace(getCellValue(row, offset+17))

		keyParts := []string{
			strings.ToLower(nomeCodigo),
			strings.ToLower(numeroArtigo),
			strings.ToLower(artigos),
		}
		key := strings.Join(keyParts, "|")

		deterministicID := ""
		if strings.TrimSpace(key) != "" {
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

		item := &model.VadeMecumOAB{
			ID:            id,
			IDTipo:        getCellValue(row, offset+0),
			Tipo:          getCellValue(row, offset+1),
			NomeCodigo:    nomeCodigo,
			Cabecalho:     getCellValue(row, offset+3),
			Titulo:        getCellValue(row, offset+4),
			TituloTexto:   getCellValue(row, offset+5),
			TituloLabel:   getCellValue(row, offset+6),
			Capitulo:      getCellValue(row, offset+7),
			CapituloTexto: getCellValue(row, offset+8),
			CapituloLabel: getCellValue(row, offset+9),
			Secao:         getCellValue(row, offset+10),
			SecaoTexto:    getCellValue(row, offset+11),
			SecaoLabel:    getCellValue(row, offset+12),
			Subsecao:      getCellValue(row, offset+13),
			SubsecaoTexto: getCellValue(row, offset+14),
			SubsecaoLabel: getCellValue(row, offset+15),
			NumeroArtigo:  numeroArtigo,
			Artigos:       artigos,
		}

		batch = append(batch, item)
	}

	if len(batch) == 0 {
		return 0, errors.New("nenhuma linha válida encontrada na planilha")
	}

	unique := deduplicateOAB(batch)

	if err := s.repo.Upsert(unique); err != nil {
		return 0, fmt.Errorf("falha ao salvar registros: %w", err)
	}

	return len(unique), nil
}

func (s *VadeMecumOABService) GetAll() ([]model.VadeMecumOAB, error) {
	return s.repo.GetAll()
}

func (s *VadeMecumOABService) GetNomeCodigoGrouped() ([]string, error) {
	return s.repo.GetNomeCodigoGrouped()
}

func (s *VadeMecumOABService) GetByID(id string) (*model.VadeMecumOAB, error) {
	trimmed := strings.TrimSpace(id)
	if trimmed == "" {
		return nil, errors.New("id inválido")
	}
	return s.repo.GetByID(trimmed)
}

func (s *VadeMecumOABService) Create(req *model.CreateVadeMecumOABRequest) (*model.VadeMecumOAB, error) {
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

	item := &model.VadeMecumOAB{
		ID:            id,
		IDTipo:        strings.TrimSpace(req.IDTipo),
		Tipo:          strings.TrimSpace(req.Tipo),
		NomeCodigo:    nome,
		Cabecalho:     strings.TrimSpace(req.Cabecalho),
		Titulo:        strings.TrimSpace(req.Titulo),
		TituloTexto:   strings.TrimSpace(req.TituloTexto),
		TituloLabel:   strings.TrimSpace(req.TituloLabel),
		Capitulo:      strings.TrimSpace(req.Capitulo),
		CapituloTexto: strings.TrimSpace(req.CapituloTexto),
		CapituloLabel: strings.TrimSpace(req.CapituloLabel),
		Secao:         strings.TrimSpace(req.Secao),
		SecaoTexto:    strings.TrimSpace(req.SecaoTexto),
		SecaoLabel:    strings.TrimSpace(req.SecaoLabel),
		Subsecao:      strings.TrimSpace(req.Subsecao),
		SubsecaoTexto: strings.TrimSpace(req.SubsecaoTexto),
		SubsecaoLabel: strings.TrimSpace(req.SubsecaoLabel),
		NumeroArtigo:  strings.TrimSpace(req.NumeroArtigo),
		Artigos:       strings.TrimSpace(req.Artigos),
	}

	if err := s.repo.Create(item); err != nil {
		return nil, err
	}

	return item, nil
}

func (s *VadeMecumOABService) Update(id string, req *model.UpdateVadeMecumOABRequest) (*model.VadeMecumOAB, error) {
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
	if req.Titulo != nil {
		existing.Titulo = strings.TrimSpace(*req.Titulo)
	}
	if req.TituloTexto != nil {
		existing.TituloTexto = strings.TrimSpace(*req.TituloTexto)
	}
	if req.TituloLabel != nil {
		existing.TituloLabel = strings.TrimSpace(*req.TituloLabel)
	}
	if req.Capitulo != nil {
		existing.Capitulo = strings.TrimSpace(*req.Capitulo)
	}
	if req.CapituloTexto != nil {
		existing.CapituloTexto = strings.TrimSpace(*req.CapituloTexto)
	}
	if req.CapituloLabel != nil {
		existing.CapituloLabel = strings.TrimSpace(*req.CapituloLabel)
	}
	if req.Secao != nil {
		existing.Secao = strings.TrimSpace(*req.Secao)
	}
	if req.SecaoTexto != nil {
		existing.SecaoTexto = strings.TrimSpace(*req.SecaoTexto)
	}
	if req.SecaoLabel != nil {
		existing.SecaoLabel = strings.TrimSpace(*req.SecaoLabel)
	}
	if req.Subsecao != nil {
		existing.Subsecao = strings.TrimSpace(*req.Subsecao)
	}
	if req.SubsecaoTexto != nil {
		existing.SubsecaoTexto = strings.TrimSpace(*req.SubsecaoTexto)
	}
	if req.SubsecaoLabel != nil {
		existing.SubsecaoLabel = strings.TrimSpace(*req.SubsecaoLabel)
	}
	if req.NumeroArtigo != nil {
		existing.NumeroArtigo = strings.TrimSpace(*req.NumeroArtigo)
	}
	if req.Artigos != nil {
		existing.Artigos = strings.TrimSpace(*req.Artigos)
	}

	if err := s.repo.Update(existing); err != nil {
		return nil, err
	}

	return existing, nil
}

func (s *VadeMecumOABService) Delete(id string) error {
	trimmed := strings.TrimSpace(id)
	if trimmed == "" {
		return errors.New("id inválido")
	}
	return s.repo.Delete(trimmed)
}

func (s *VadeMecumOABService) Reset() error {
	return s.repo.DeleteAll()
}

func deduplicateOAB(items []*model.VadeMecumOAB) []*model.VadeMecumOAB {
	index := make(map[string]int, len(items))
	unique := make([]*model.VadeMecumOAB, 0, len(items))

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


