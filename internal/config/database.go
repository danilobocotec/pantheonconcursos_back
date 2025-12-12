package config

import (
	"errors"
	"log"

	"github.com/thepantheon/api/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *Config) (*gorm.DB, error) {
	// Fallback para ambiente de desenvolvimento local: força localhost e desabilita SSL
	if cfg.Server.Env == "development" {
		cfg.Database.Host = "localhost"
		cfg.Database.SSLMode = "disable"
		// Mantém o nome do banco vindo do .env (ex: "postgres")
	}

	dsn := cfg.Database.GetDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Database connection established successfully")
	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&model.Plan{},
		&model.User{},
		// Add more models here as needed
	); err != nil {
		return err
	}

	// Ensure legacy column "name" is renamed to "full_name"
	migrator := db.Migrator()
	if migrator.HasColumn(&model.User{}, "name") {
		if migrator.HasColumn(&model.User{}, "full_name") {
			// Copy existing data before dropping legacy column
			if err := db.Exec(`UPDATE users SET full_name = name WHERE full_name IS NULL OR full_name = ''`).Error; err != nil {
				return err
			}
			if err := migrator.DropColumn(&model.User{}, "name"); err != nil {
				return err
			}
		} else {
			if err := migrator.RenameColumn(&model.User{}, "name", "full_name"); err != nil {
				return err
			}
		}
	}

	// Seed default plans if they don't exist
	defaultPlans := []model.Plan{
		{
			Name:        "Plano 1 — Fase 1 (Anual)",
			Description: "Acesso anual completo à 1ª fase da OAB.",
			Phase:       "fase_1",
			Duration:    "anual",
			Active:      true,
		},
		{
			Name:        "Plano 2 — Fase 1 (Vitalício)",
			Description: "Acesso vitalício à 1ª fase da OAB.",
			Phase:       "fase_1",
			Duration:    "vitalicio",
			Active:      true,
		},
		{
			Name:        "Plano 3 — Fase 1 + Fase 2 (Anual)",
			Description: "Acesso anual à 1ª e 2ª fases da OAB.",
			Phase:       "fase_1_2",
			Duration:    "anual",
			Active:      true,
		},
	}

	for _, plan := range defaultPlans {
		var existing model.Plan
		if err := db.Where("name = ?", plan.Name).First(&existing).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err := db.Create(&plan).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}

	return nil
}
