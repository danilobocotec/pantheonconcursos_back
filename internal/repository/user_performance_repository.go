package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/thepantheon/api/internal/model"
	"gorm.io/gorm"
)

type UserPerformanceRepository struct {
	db *gorm.DB
}

func NewUserPerformanceRepository(db *gorm.DB) *UserPerformanceRepository {
	return &UserPerformanceRepository{db: db}
}

func (r *UserPerformanceRepository) Create(item *model.UserPerformance) error {
	return r.db.Create(item).Error
}

func (r *UserPerformanceRepository) GetByUser(userID uuid.UUID, startDate, endDate *time.Time) ([]model.UserPerformance, error) {
	var items []model.UserPerformance
	query := r.db.Where("user_id = ?", userID)
	if startDate != nil {
		query = query.Where("recorded_at >= ?", *startDate)
	}
	if endDate != nil {
		query = query.Where("recorded_at <= ?", *endDate)
	}
	if err := query.Order("recorded_at ASC").Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *UserPerformanceRepository) GetTotals(userID uuid.UUID, startDate, endDate *time.Time) (*model.UserPerformanceSummary, error) {
	var totals struct {
		TotalQuestions   int
		CorrectQuestions int
		WrongQuestions   int
	}
	query := r.db.Model(&model.UserPerformance{}).Where("user_id = ?", userID)
	if startDate != nil {
		query = query.Where("recorded_at >= ?", *startDate)
	}
	if endDate != nil {
		query = query.Where("recorded_at <= ?", *endDate)
	}
	if err := query.Select(
		"COALESCE(SUM(total_questions), 0) AS total_questions",
		"COALESCE(SUM(correct_questions), 0) AS correct_questions",
		"COALESCE(SUM(wrong_questions), 0) AS wrong_questions",
	).Scan(&totals).Error; err != nil {
		return nil, err
	}
	return &model.UserPerformanceSummary{
		TotalQuestions:   totals.TotalQuestions,
		CorrectQuestions: totals.CorrectQuestions,
		WrongQuestions:   totals.WrongQuestions,
	}, nil
}
