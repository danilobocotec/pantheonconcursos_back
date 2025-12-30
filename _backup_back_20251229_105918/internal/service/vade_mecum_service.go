package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/thepantheon/api/internal/model"
	"github.com/thepantheon/api/internal/repository"
)

type VadeMecumService struct {
	repo *repository.VadeMecumRepository
}

var allowedCategories = []string{
	"constituicao",
	"codigos",
	"leis",
	"jurisprudencia",
	"oab",
	"estatutos",
}

func NewVadeMecumService(repo *repository.VadeMecumRepository) *VadeMecumService {
	return &VadeMecumService{repo: repo}
}

func (s *VadeMecumService) validateCategory(category string) bool {
	for _, c := range allowedCategories {
		if c == category {
			return true
		}
	}
	return false
}

func (s *VadeMecumService) Create(req *model.CreateVadeMecumRequest) (*model.VadeMecum, error) {
	if !s.validateCategory(req.Category) {
		return nil, errors.New("categoria inválida")
	}

	vm := &model.VadeMecum{
		Title:       req.Title,
		Description: req.Description,
		Content:     req.Content,
		FileURL:     req.FileURL,
		Category:    req.Category,
		Header:      req.Header,
		TitleID:     req.TitleID,
		TitleName:   req.TitleName,
		TitleText:   req.TitleText,
		ChapterID:   req.ChapterID,
		ChapterName: req.ChapterName,
		ChapterText: req.ChapterText,
	}
	if err := s.repo.Create(vm); err != nil {
		return nil, err
	}
	return vm, nil
}

func (s *VadeMecumService) GetAll() ([]model.VadeMecum, error) {
	return s.repo.GetAll()
}

func (s *VadeMecumService) GetByCategory(category string) ([]model.VadeMecum, error) {
	if !s.validateCategory(category) {
		return nil, errors.New("categoria inválida")
	}
	return s.repo.GetByCategory(category)
}

func (s *VadeMecumService) GetByID(id uuid.UUID) (*model.VadeMecum, error) {
	return s.repo.GetByID(id)
}

func (s *VadeMecumService) Update(id uuid.UUID, req *model.UpdateVadeMecumRequest) (*model.VadeMecum, error) {
	vm, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Title != "" {
		vm.Title = req.Title
	}
	if req.Description != "" {
		vm.Description = req.Description
	}
	if req.Content != "" {
		vm.Content = req.Content
	}
	if req.FileURL != "" {
		vm.FileURL = req.FileURL
	}
	if req.Category != "" {
		if !s.validateCategory(req.Category) {
			return nil, errors.New("categoria inválida")
		}
		vm.Category = req.Category
	}
	if req.Header != "" {
		vm.Header = req.Header
	}
	if req.TitleID != "" {
		vm.TitleID = req.TitleID
	}
	if req.TitleName != "" {
		vm.TitleName = req.TitleName
	}
	if req.TitleText != "" {
		vm.TitleText = req.TitleText
	}
	if req.ChapterID != "" {
		vm.ChapterID = req.ChapterID
	}
	if req.ChapterName != "" {
		vm.ChapterName = req.ChapterName
	}
	if req.ChapterText != "" {
		vm.ChapterText = req.ChapterText
	}

	if err := s.repo.Update(vm); err != nil {
		return nil, err
	}

	return vm, nil
}

func (s *VadeMecumService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
