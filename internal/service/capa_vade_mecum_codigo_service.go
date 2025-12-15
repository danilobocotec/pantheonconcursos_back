package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/thepantheon/api/internal/model"
	"github.com/thepantheon/api/internal/repository"
	"gorm.io/gorm"
)

type CapaVadeMecumCodigoService struct {
	repo *repository.CapaVadeMecumCodigoRepository
}

func NewCapaVadeMecumCodigoService(repo *repository.CapaVadeMecumCodigoRepository) *CapaVadeMecumCodigoService {
	return &CapaVadeMecumCodigoService{repo: repo}
}

func (s *CapaVadeMecumCodigoService) GetAll() ([]model.CapaVadeMecumCodigo, error) {
	return s.repo.GetAll()
}

func (s *CapaVadeMecumCodigoService) GetByNomeCodigo(nomecodigo string) (*model.CapaVadeMecumCodigo, error) {
	trimmed := strings.TrimSpace(nomecodigo)
	if trimmed == "" {
		return nil, errors.New("nomecodigo inválido")
	}
	return s.repo.GetByNomeCodigo(trimmed)
}

func (s *CapaVadeMecumCodigoService) Create(req *model.CreateCapaVadeMecumCodigoRequest) (*model.CapaVadeMecumCodigo, error) {
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

	item := &model.CapaVadeMecumCodigo{
		NomeCodigo: nome,
		Cabecalho:  strings.TrimSpace(req.Cabecalho),
		Grupo:      strings.TrimSpace(req.Grupo),
	}

	if err := s.repo.Create(item); err != nil {
		return nil, err
	}

	return item, nil
}

func (s *CapaVadeMecumCodigoService) Update(nomecodigo string, req *model.UpdateCapaVadeMecumCodigoRequest) (*model.CapaVadeMecumCodigo, error) {
	if req == nil {
		return nil, errors.New("payload obrigatório")
	}

	nome := strings.TrimSpace(nomecodigo)
	if nome == "" {
		return nil, errors.New("nomecodigo inválido")
	}

	existing, err := s.repo.GetByNomeCodigo(nome)
	if err != nil {
		return nil, err
	}

	if req.Cabecalho != nil {
		existing.Cabecalho = strings.TrimSpace(*req.Cabecalho)
	}
	if req.Grupo != nil {
		existing.Grupo = strings.TrimSpace(*req.Grupo)
	}

	if err := s.repo.Update(existing); err != nil {
		return nil, err
	}

	return existing, nil
}
