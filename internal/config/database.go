package config

import (
	"errors"
	"log"

	"github.com/thepantheon/api/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *Config) (*gorm.DB, error) {
	// Fallback somente quando nenhum host foi configurado explicitamente
	if cfg.Server.Env == "development" &&
		(cfg.Database.Host == "" || cfg.Database.Host == "localhost") {
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
		&model.VadeMecum{},
		&model.VadeMecumCodigo{},
		&model.User{},
		// Add more models here as needed
	); err != nil {
		return err
	}

	// Ensure legacy column "name" is renamed to "full_name"
	migrator := db.Migrator()

	// Align columns to the exact schema required for vade_mecum_codigos.
	// Supports both previous goose SQL naming and GORM default snake_case naming.
	renameColumnIfNeeded := func(table, oldName, newName string) error {
		if migrator.HasColumn(table, oldName) && !migrator.HasColumn(table, newName) {
			if err := migrator.RenameColumn(table, oldName, newName); err != nil {
				return err
			}
		}
		return nil
	}

	const vadeCodigosTable = "vade_mecum_codigos"
	if err := renameColumnIfNeeded(vadeCodigosTable, "nome_codigo", "nomecodigo"); err != nil {
		return err
	}
	if err := renameColumnIfNeeded(vadeCodigosTable, "cabecalho", "Cabecalho"); err != nil {
		return err
	}
	if err := renameColumnIfNeeded(vadeCodigosTable, "parte", "PARTE"); err != nil {
		return err
	}
	if err := renameColumnIfNeeded(vadeCodigosTable, "livro_id", "idlivro"); err != nil {
		return err
	}
	if err := renameColumnIfNeeded(vadeCodigosTable, "livro_texto", "livrotexto"); err != nil {
		return err
	}
	if err := renameColumnIfNeeded(vadeCodigosTable, "titulo_id", "idtitulo"); err != nil {
		return err
	}
	if err := renameColumnIfNeeded(vadeCodigosTable, "titulo_texto", "titulotexto"); err != nil {
		return err
	}
	if err := renameColumnIfNeeded(vadeCodigosTable, "subtitulo_id", "idsubtitulo"); err != nil {
		return err
	}
	if err := renameColumnIfNeeded(vadeCodigosTable, "subtitulo_texto", "subtitulotexto"); err != nil {
		return err
	}
	if err := renameColumnIfNeeded(vadeCodigosTable, "capitulo_id", "idcapitulo"); err != nil {
		return err
	}
	if err := renameColumnIfNeeded(vadeCodigosTable, "capitulo_texto", "capitulotexto"); err != nil {
		return err
	}
	if err := renameColumnIfNeeded(vadeCodigosTable, "secao_id", "idsecao"); err != nil {
		return err
	}
	if err := renameColumnIfNeeded(vadeCodigosTable, "secao_texto", "secaotexto"); err != nil {
		return err
	}
	if err := renameColumnIfNeeded(vadeCodigosTable, "subsecao_id", "idsubsecao"); err != nil {
		return err
	}
	if err := renameColumnIfNeeded(vadeCodigosTable, "subsecao_texto", "subsecaotexto"); err != nil {
		return err
	}
	if err := renameColumnIfNeeded(vadeCodigosTable, "numero_artigo", "num_artigo"); err != nil {
		return err
	}
	if err := renameColumnIfNeeded(vadeCodigosTable, "normativo", "Normativo"); err != nil {
		return err
	}
	if err := renameColumnIfNeeded(vadeCodigosTable, "ordem", "Ordem"); err != nil {
		return err
	}

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

	if migrator.HasColumn(&model.User{}, "role") {
		if err := db.Model(&model.User{}).
			Where("role IS NULL OR role = ''").
			Update("role", "user").Error; err != nil {
			return err
		}
	}

	if migrator.HasColumn(&model.VadeMecum{}, "category") {
		if err := db.Model(&model.VadeMecum{}).
			Where("category IS NULL OR category = ''").
			Update("category", "constituicao").Error; err != nil {
			return err
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
