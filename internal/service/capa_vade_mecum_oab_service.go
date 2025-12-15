package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/thepantheon/api/internal/model"
	"github.com/thepantheon/api/internal/repository"
	"gorm.io/gorm"
)

type CapaVadeMecumOABService struct {
	repo *repository.CapaVadeMecumOABRepository
}

func NewCapaVadeMecumOABService(repo *repository.CapaVadeMecumOABRepository) *CapaVadeMecumOABService {
	return &CapaVadeMecumOABService{repo: repo}
}

func (s *CapaVadeMecumOABService) GetAll() ([]model.CapaVadeMecumOAB, error) {
	return s.repo.GetAll()
}

func (s *CapaVadeMecumOABService) GetByID(id string) (*model.CapaVadeMecumOAB, error) {
	trimmed := strings.TrimSpace(id)
	if trimmed == "" {
		return nil, errors.New("id inválido")
	}
	return s.repo.GetByID(trimmed)
}

func (s *CapaVadeMecumOABService) GetByNomeCodigo(nomecodigo string) (*model.CapaVadeMecumOAB, error) {
	trimmed := strings.TrimSpace(nomecodigo)
	if trimmed == "" {
		return nil, errors.New("nomecodigo inválido")
	}
	return s.repo.GetByNomeCodigo(trimmed)
}

func (s *CapaVadeMecumOABService) Create(req *model.CreateCapaVadeMecumOABRequest) (*model.CapaVadeMecumOAB, error) {
	if req == nil {
		return nil, errors.New("payload obrigatório")
	}

	nome := strings.TrimSpace(req.NomeCodigo)
	if nome == "" {
		return nil, errors.New("nomecodigo é obrigatório")
	}

	if _, err := s.repo.GetByNomeCodigo(nome); err == nil {
		return nil, fmt.Errorf("capa '%s' já existe", nome)
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	grupo := strings.TrimSpace(req.Grupo)
	if grupo != "" {
		if existing, err := s.repo.GetByGrupo(grupo); err == nil {
			return nil, fmt.Errorf("grupo '%s' já está vinculado à capa '%s'", grupo, existing.NomeCodigo)
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	item := &model.CapaVadeMecumOAB{
		NomeCodigo: nome,
		Cabecalho:  strings.TrimSpace(req.Cabecalho),
		Grupo:      grupo,
	}

	if err := s.repo.Create(item); err != nil {
		return nil, err
	}

	return item, nil
}

func (s *CapaVadeMecumOABService) Update(id string, req *model.UpdateCapaVadeMecumOABRequest) (*model.CapaVadeMecumOAB, error) {
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

	if req.Cabecalho != nil {
		existing.Cabecalho = strings.TrimSpace(*req.Cabecalho)
	}
	if req.Grupo != nil {
		grupo := strings.TrimSpace(*req.Grupo)
		if grupo != "" {
			if found, err := s.repo.GetByGrupo(grupo); err == nil && found.ID != existing.ID {
				return nil, fmt.Errorf("grupo '%s' já está vinculado à capa '%s'", grupo, found.NomeCodigo)
			} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, err
			}
		}
		existing.Grupo = grupo
	}

	if err := s.repo.Update(existing); err != nil {
		return nil, err
	}

	return existing, nil
}
