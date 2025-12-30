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

type VadeMecumLeiService struct {
	repo *repository.VadeMecumLeiRepository
}

func NewVadeMecumLeiService(repo *repository.VadeMecumLeiRepository) *VadeMecumLeiService {
	return &VadeMecumLeiService{repo: repo}
}

var vadeMecumLeisHeaders = []string{
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

		item := &model.VadeMecumLei{
			IDTipo:        getCellValue(row, 0),
			Tipo:          getCellValue(row, 1),
			NomeCodigo:    getCellValue(row, 2),
			Cabecalho:     getCellValue(row, 3),
			IDParte:       getCellValue(row, 4),
			Parte:         getCellValue(row, 5),
			ParteTexto:    getCellValue(row, 6),
			IDTitulo:      getCellValue(row, 7),
			Titulo:        getCellValue(row, 8),
			TituloTexto:   getCellValue(row, 9),
			IDCapitulo:    getCellValue(row, 10),
			Capitulo:      getCellValue(row, 11),
			CapituloTexto: getCellValue(row, 12),
			IDSecao:       getCellValue(row, 13),
			Secao:         getCellValue(row, 14),
			SecaoTexto:    getCellValue(row, 15),
			IDSubsecao:    getCellValue(row, 16),
			Subsecao:      getCellValue(row, 17),
			SubsecaoTexto: getCellValue(row, 18),
			NumeroArtigo:  getCellValue(row, 19),
			Artigos:       getCellValue(row, 20),
			Ordem:         getCellValue(row, 21),
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

func (s *VadeMecumLeiService) Create(req *model.CreateVadeMecumLeiRequest) (*model.VadeMecumLei, error) {
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
			return nil, fmt.Errorf("lei com id '%s' já existe", id)
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	item := &model.VadeMecumLei{
		ID:            id,
		IDTipo:        strings.TrimSpace(req.IDTipo),
		Tipo:          strings.TrimSpace(req.Tipo),
		NomeCodigo:    nome,
		Cabecalho:     strings.TrimSpace(req.Cabecalho),
		IDParte:       strings.TrimSpace(req.IDParte),
		Parte:         strings.TrimSpace(req.Parte),
		ParteTexto:    strings.TrimSpace(req.ParteTexto),
		IDTitulo:      strings.TrimSpace(req.IDTitulo),
		Titulo:        strings.TrimSpace(req.Titulo),
		TituloTexto:   strings.TrimSpace(req.TituloTexto),
		IDCapitulo:    strings.TrimSpace(req.IDCapitulo),
		Capitulo:      strings.TrimSpace(req.Capitulo),
		CapituloTexto: strings.TrimSpace(req.CapituloTexto),
		IDSecao:       strings.TrimSpace(req.IDSecao),
		Secao:         strings.TrimSpace(req.Secao),
		SecaoTexto:    strings.TrimSpace(req.SecaoTexto),
		IDSubsecao:    strings.TrimSpace(req.IDSubsecao),
		Subsecao:      strings.TrimSpace(req.Subsecao),
		SubsecaoTexto: strings.TrimSpace(req.SubsecaoTexto),
		NumeroArtigo:  strings.TrimSpace(req.NumeroArtigo),
		Artigos:       strings.TrimSpace(req.Artigos),
		Ordem:         strings.TrimSpace(req.Ordem),
	}

	if err := s.repo.Create(item); err != nil {
		return nil, err
	}

	return item, nil
}

func (s *VadeMecumLeiService) GetAll() ([]model.VadeMecumLei, error) {
	return s.repo.GetAll()
}

func (s *VadeMecumLeiService) GrupoServico() ([]model.VadeMecumLeiGrupoServico, error) {
	return s.repo.GetGrupoServico()
}

func (s *VadeMecumLeiService) GetByID(id string) (*model.VadeMecumLei, error) {
	trimmed := strings.TrimSpace(id)
	if trimmed == "" {
		return nil, errors.New("id inválido")
	}
	return s.repo.GetByID(trimmed)
}

func (s *VadeMecumLeiService) Update(id string, req *model.UpdateVadeMecumLeiRequest) (*model.VadeMecumLei, error) {
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
	if req.IDParte != nil {
		existing.IDParte = strings.TrimSpace(*req.IDParte)
	}
	if req.Parte != nil {
		existing.Parte = strings.TrimSpace(*req.Parte)
	}
	if req.ParteTexto != nil {
		existing.ParteTexto = strings.TrimSpace(*req.ParteTexto)
	}
	if req.IDTitulo != nil {
		existing.IDTitulo = strings.TrimSpace(*req.IDTitulo)
	}
	if req.Titulo != nil {
		existing.Titulo = strings.TrimSpace(*req.Titulo)
	}
	if req.TituloTexto != nil {
		existing.TituloTexto = strings.TrimSpace(*req.TituloTexto)
	}
	if req.IDCapitulo != nil {
		existing.IDCapitulo = strings.TrimSpace(*req.IDCapitulo)
	}
	if req.Capitulo != nil {
		existing.Capitulo = strings.TrimSpace(*req.Capitulo)
	}
	if req.CapituloTexto != nil {
		existing.CapituloTexto = strings.TrimSpace(*req.CapituloTexto)
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
	if req.Artigos != nil {
		existing.Artigos = strings.TrimSpace(*req.Artigos)
	}
	if req.Ordem != nil {
		existing.Ordem = strings.TrimSpace(*req.Ordem)
	}

	if err := s.repo.Update(existing); err != nil {
		return nil, err
	}

	return existing, nil
}

func (s *VadeMecumLeiService) Delete(id string) error {
	trimmed := strings.TrimSpace(id)
	if trimmed == "" {
		return errors.New("id inválido")
	}
	return s.repo.Delete(trimmed)
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
