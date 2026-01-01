package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserPerformance struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	UserID           uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	TotalQuestions   int       `gorm:"not null" json:"total_questoes"`
	CorrectQuestions int       `gorm:"not null" json:"questoes_corretas"`
	WrongQuestions   int       `gorm:"not null" json:"questoes_erradas"`
	AccuracyPercent  float64   `gorm:"not null" json:"percentual_acerto"`
	RecordedAt       time.Time `gorm:"not null" json:"data_gravacao"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func (u *UserPerformance) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	if u.RecordedAt.IsZero() {
		u.RecordedAt = time.Now()
	}
	return nil
}

type CreateUserPerformanceRequest struct {
	TotalQuestoes     int        `json:"total_questoes" binding:"min=0"`
	QuestoesCorretas  int        `json:"questoes_corretas" binding:"min=0"`
	QuestoesErradas   int        `json:"questoes_erradas" binding:"min=0"`
	DataGravacao      *time.Time `json:"data_gravacao"`
}

type UserPerformanceSummary struct {
	TotalQuestions   int     `json:"total_questoes"`
	CorrectQuestions int     `json:"questoes_corretas"`
	WrongQuestions   int     `json:"questoes_erradas"`
	AccuracyPercent  float64 `json:"percentual_acerto"`
}
