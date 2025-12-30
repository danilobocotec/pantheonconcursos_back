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

type VadeMecumConstituicaoService struct {
	repo *repository.VadeMecumConstituicaoRepository
}

func NewVadeMecumConstituicaoService(repo *repository.VadeMecumConstituicaoRepository) *VadeMecumConstituicaoService {
	return &VadeMecumConstituicaoService{repo: repo}
}

var constituicaoHeaders = []string{
	"idtipo",
	"tipo",
	"cabecalho",
	"idtitulo",
	"titulo",
	"textodotitulo",
	"idcapitulo",
	"capítulo",
	"textocapítulo",
	"idsecao",
	"secao",
	"texttosecao",
	"idsubsecao",
	"subsecao",
	"subsecaotexto",
	"Normativo",
}

func normalizeConstituicaoHeader(row []string) []string {
	normalized := make([]string, len(row))
	for i, value := range row {
		normalized[i] = strings.TrimSpace(value)
	}
	return normalized
}

func headersEqual(header []string, expected []string) bool {
	if len(header) < len(expected) {
		return false
	}

	for idx, exp := range expected {
		if header[idx] != exp {
			return false
		}
	}
	return true
}

func (s *VadeMecumConstituicaoService) ImportConstituicao(r io.Reader) (int, error) {
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

	header := normalizeConstituicaoHeader(rows[0])
	if !headersEqual(header, constituicaoHeaders) {
		return 0, fmt.Errorf("cabeçalho inválido: esperado %v", constituicaoHeaders)
	}

	var batch []*model.VadeMecumConstituicao

	for idx, row := range rows[1:] {
		if isRowEmpty(row) {
			continue
		}

		normativo := getCellValue(row, 15)
		registroID := strings.TrimSpace(normativo)
		if registroID != "" {
			registroID = uuid.NewSHA1(uuid.NameSpaceOID, []byte(registroID)).String()
		} else {
			registroID = uuid.NewString()
		}

		item := &model.VadeMecumConstituicao{
			RegistroID:    registroID,
			IDTipo:        getCellValue(row, 0),
			Tipo:          getCellValue(row, 1),
			Cabecalho:     getCellValue(row, 2),
			IDTitulo:      getCellValue(row, 3),
			Titulo:        getCellValue(row, 4),
			TextoDoTitulo: getCellValue(row, 5),
			IDCapitulo:    getCellValue(row, 6),
			Capitulo:      getCellValue(row, 7),
			TextoCapitulo: getCellValue(row, 8),
			IDSecao:       getCellValue(row, 9),
			Secao:         getCellValue(row, 10),
			TextoSecao:    getCellValue(row, 11),
			IDSubsecao:    getCellValue(row, 12),
			Subsecao:      getCellValue(row, 13),
			TextoSubsecao: getCellValue(row, 14),
			Normativo:     normativo,
		}

		if strings.TrimSpace(item.Normativo) == "" {
			return 0, fmt.Errorf("linha %d: Normativo é obrigatório", idx+2)
		}

		batch = append(batch, item)
	}

	if len(batch) == 0 {
		return 0, errors.New("nenhuma linha válida encontrada na planilha")
	}

	if err := s.repo.Upsert(batch); err != nil {
		return 0, fmt.Errorf("falha ao salvar registros: %w", err)
	}

	return len(batch), nil
}

func (s *VadeMecumConstituicaoService) GetAll() ([]model.VadeMecumConstituicao, error) {
	return s.repo.GetAll()
}

func (s *VadeMecumConstituicaoService) GrupoServico() ([]model.VadeMecumConstituicaoGrupoServico, error) {
	return s.repo.GetGrupoServico()
}

func (s *VadeMecumConstituicaoService) GetByID(id uuid.UUID) (*model.VadeMecumConstituicao, error) {
	return s.repo.GetByID(id)
}

func (s *VadeMecumConstituicaoService) Create(req *model.CreateVadeMecumConstituicaoRequest) (*model.VadeMecumConstituicao, error) {
	if req == nil {
		return nil, errors.New("payload obrigatório")
	}

	normativo := strings.TrimSpace(req.Normativo)
	if normativo == "" {
		return nil, errors.New("normativo é obrigatório")
	}

	registroID := strings.TrimSpace(req.RegistroID)
	if registroID == "" {
		registroID = uuid.NewSHA1(uuid.NameSpaceOID, []byte(normativo)).String()
	}

	item := &model.VadeMecumConstituicao{
		RegistroID:    registroID,
		IDTipo:        strings.TrimSpace(req.IDTipo),
		Tipo:          strings.TrimSpace(req.Tipo),
		Cabecalho:     strings.TrimSpace(req.Cabecalho),
		IDTitulo:      strings.TrimSpace(req.IDTitulo),
		Titulo:        strings.TrimSpace(req.Titulo),
		TextoDoTitulo: strings.TrimSpace(req.TextoDoTitulo),
		IDCapitulo:    strings.TrimSpace(req.IDCapitulo),
		Capitulo:      strings.TrimSpace(req.Capitulo),
		TextoCapitulo: strings.TrimSpace(req.TextoCapitulo),
		IDSecao:       strings.TrimSpace(req.IDSecao),
		Secao:         strings.TrimSpace(req.Secao),
		TextoSecao:    strings.TrimSpace(req.TextoSecao),
		IDSubsecao:    strings.TrimSpace(req.IDSubsecao),
		Subsecao:      strings.TrimSpace(req.Subsecao),
		TextoSubsecao: strings.TrimSpace(req.TextoSubsecao),
		Normativo:     normativo,
	}

	if err := s.repo.Create(item); err != nil {
		return nil, err
	}

	return item, nil
}

func (s *VadeMecumConstituicaoService) Update(id uuid.UUID, req *model.UpdateVadeMecumConstituicaoRequest) (*model.VadeMecumConstituicao, error) {
	if req == nil {
		return nil, errors.New("payload obrigatório")
	}

	item, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.RegistroID != nil {
		item.RegistroID = strings.TrimSpace(*req.RegistroID)
	}
	if req.IDTipo != nil {
		item.IDTipo = strings.TrimSpace(*req.IDTipo)
	}
	if req.Tipo != nil {
		item.Tipo = strings.TrimSpace(*req.Tipo)
	}
	if req.Cabecalho != nil {
		item.Cabecalho = strings.TrimSpace(*req.Cabecalho)
	}
	if req.IDTitulo != nil {
		item.IDTitulo = strings.TrimSpace(*req.IDTitulo)
	}
	if req.Titulo != nil {
		item.Titulo = strings.TrimSpace(*req.Titulo)
	}
	if req.TextoDoTitulo != nil {
		item.TextoDoTitulo = strings.TrimSpace(*req.TextoDoTitulo)
	}
	if req.IDCapitulo != nil {
		item.IDCapitulo = strings.TrimSpace(*req.IDCapitulo)
	}
	if req.Capitulo != nil {
		item.Capitulo = strings.TrimSpace(*req.Capitulo)
	}
	if req.TextoCapitulo != nil {
		item.TextoCapitulo = strings.TrimSpace(*req.TextoCapitulo)
	}
	if req.IDSecao != nil {
		item.IDSecao = strings.TrimSpace(*req.IDSecao)
	}
	if req.Secao != nil {
		item.Secao = strings.TrimSpace(*req.Secao)
	}
	if req.TextoSecao != nil {
		item.TextoSecao = strings.TrimSpace(*req.TextoSecao)
	}
	if req.IDSubsecao != nil {
		item.IDSubsecao = strings.TrimSpace(*req.IDSubsecao)
	}
	if req.Subsecao != nil {
		item.Subsecao = strings.TrimSpace(*req.Subsecao)
	}
	if req.TextoSubsecao != nil {
		item.TextoSubsecao = strings.TrimSpace(*req.TextoSubsecao)
	}
	if req.Normativo != nil {
		item.Normativo = strings.TrimSpace(*req.Normativo)
	}

	if err := s.repo.Update(item); err != nil {
		return nil, err
	}

	return item, nil
}

func (s *VadeMecumConstituicaoService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
