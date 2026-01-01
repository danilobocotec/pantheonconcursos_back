package service

import (
	"errors"
	"math"
	"time"

	"github.com/google/uuid"
	"github.com/thepantheon/api/internal/model"
	"github.com/thepantheon/api/internal/repository"
)

type UserPerformanceService struct {
	repo *repository.UserPerformanceRepository
}

func NewUserPerformanceService(repo *repository.UserPerformanceRepository) *UserPerformanceService {
	return &UserPerformanceService{repo: repo}
}

func (s *UserPerformanceService) Create(userID uuid.UUID, req *model.CreateUserPerformanceRequest) (*model.UserPerformance, error) {
	if req == nil {
		return nil, errors.New("payload obrigatorio")
	}
	if userID == uuid.Nil {
		return nil, errors.New("usuario invalido")
	}
	if req.TotalQuestoes < 0 || req.QuestoesCorretas < 0 || req.QuestoesErradas < 0 {
		return nil, errors.New("valores invalidos")
	}
	totalRealizado := req.QuestoesCorretas + req.QuestoesErradas
	if req.TotalQuestoes != totalRealizado {
		return nil, errors.New("total de questoes deve ser igual a corretas + erradas")
	}

	recordedAt := time.Now()
	if req.DataGravacao != nil && !req.DataGravacao.IsZero() {
		recordedAt = *req.DataGravacao
	}

	accuracy := 0.0
	if req.TotalQuestoes > 0 {
		accuracy = (float64(req.QuestoesCorretas) / float64(req.TotalQuestoes)) * 100
		accuracy = math.Round(accuracy*100) / 100
	}

	item := &model.UserPerformance{
		UserID:           userID,
		TotalQuestions:   req.TotalQuestoes,
		CorrectQuestions: req.QuestoesCorretas,
		WrongQuestions:   req.QuestoesErradas,
		AccuracyPercent:  accuracy,
		RecordedAt:       recordedAt,
	}

	if err := s.repo.Create(item); err != nil {
		return nil, err
	}

	return item, nil
}

func (s *UserPerformanceService) GetByUser(userID uuid.UUID, startDate, endDate *time.Time) ([]model.UserPerformance, error) {
	if userID == uuid.Nil {
		return nil, errors.New("usuario invalido")
	}
	return s.repo.GetByUser(userID, startDate, endDate)
}

func (s *UserPerformanceService) GetSummary(userID uuid.UUID, startDate, endDate *time.Time) (*model.UserPerformanceSummary, error) {
	if userID == uuid.Nil {
		return nil, errors.New("usuario invalido")
	}
	totals, err := s.repo.GetTotals(userID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	if totals.TotalQuestions > 0 {
		totals.AccuracyPercent = (float64(totals.CorrectQuestions) / float64(totals.TotalQuestions)) * 100
		totals.AccuracyPercent = math.Round(totals.AccuracyPercent*100) / 100
	}
	return totals, nil
}
