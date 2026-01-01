package config

import (
	"errors"
	"log"
	"strings"

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
		&model.Questao{},
		&model.VadeMecum{},
		&model.VadeMecumCodigo{},
		&model.VadeMecumEstatuto{},
		&model.VadeMecumConstituicao{},
		&model.VadeMecumLei{},
		&model.VadeMecumOAB{},
		&model.CapaVadeMecumCodigo{},
		&model.CapaVadeMecumOAB{},
		&model.CapaVadeMecumJurisprudencia{},
		&model.VadeMecumJurisprudencia{},
		&model.Course{},
		&model.CourseCategory{},
		&model.CourseModule{},
		&model.CourseItem{},
		&model.CourseModuleItem{},
		&model.UserPerformance{},
		&model.User{},
		// Add more models here as needed
	); err != nil {
		return err
	}

	// Guarantee all codigo columns stay as TEXT to support planilhas sem limite de caracteres.
	if err := db.Exec(`ALTER TABLE vade_mecum_codigos
		ALTER COLUMN idtipo TYPE TEXT,
		ALTER COLUMN tipo TYPE TEXT,
		ALTER COLUMN idcodigo TYPE TEXT,
		ALTER COLUMN nomecodigo TYPE TEXT,
		ALTER COLUMN "Cabecalho" TYPE TEXT,
		ALTER COLUMN "PARTE" TYPE TEXT,
		ALTER COLUMN idlivro TYPE TEXT,
		ALTER COLUMN livro TYPE TEXT,
		ALTER COLUMN livrotexto TYPE TEXT,
		ALTER COLUMN idtitulo TYPE TEXT,
		ALTER COLUMN titulo TYPE TEXT,
		ALTER COLUMN titulotexto TYPE TEXT,
		ALTER COLUMN idsubtitulo TYPE TEXT,
		ALTER COLUMN subtitulo TYPE TEXT,
		ALTER COLUMN subtitulotexto TYPE TEXT,
		ALTER COLUMN idcapitulo TYPE TEXT,
		ALTER COLUMN capitulo TYPE TEXT,
		ALTER COLUMN capitulotexto TYPE TEXT,
		ALTER COLUMN idsecao TYPE TEXT,
		ALTER COLUMN secao TYPE TEXT,
		ALTER COLUMN secaotexto TYPE TEXT,
		ALTER COLUMN idsubsecao TYPE TEXT,
		ALTER COLUMN subsecao TYPE TEXT,
		ALTER COLUMN subsecaotexto TYPE TEXT,
		ALTER COLUMN num_artigo TYPE TEXT,
		ALTER COLUMN "Normativo" TYPE TEXT,
		ALTER COLUMN "Ordem" TYPE TEXT`).Error; err != nil {
		return err
	}

	if err := db.Exec(`ALTER TABLE vade_mecum_leis
		ALTER COLUMN id TYPE TEXT,
		ALTER COLUMN idtipo TYPE TEXT,
		ALTER COLUMN tipo TYPE TEXT,
		ALTER COLUMN nomecodigo TYPE TEXT,
		ALTER COLUMN "Cabecalho" TYPE TEXT,
		ALTER COLUMN "idPARTE" TYPE TEXT,
		ALTER COLUMN "PARTE" TYPE TEXT,
		ALTER COLUMN "PARTETEXTO" TYPE TEXT,
		ALTER COLUMN idtitulo TYPE TEXT,
		ALTER COLUMN titulo TYPE TEXT,
		ALTER COLUMN titulotexto TYPE TEXT,
		ALTER COLUMN idcapitulo TYPE TEXT,
		ALTER COLUMN capitulo TYPE TEXT,
		ALTER COLUMN capitulotexto TYPE TEXT,
		ALTER COLUMN idsecao TYPE TEXT,
		ALTER COLUMN secao TYPE TEXT,
		ALTER COLUMN secaotexto TYPE TEXT,
		ALTER COLUMN idsubsecao TYPE TEXT,
		ALTER COLUMN subsecao TYPE TEXT,
		ALTER COLUMN subsecaotexto TYPE TEXT,
		ALTER COLUMN num_artigo TYPE TEXT,
		ALTER COLUMN "Artigos" TYPE TEXT,
		ALTER COLUMN "Ordem" TYPE TEXT`).Error; err != nil {
		return err
	}

	if err := db.Exec(`ALTER TABLE vade_mecum_oab
		ALTER COLUMN id TYPE TEXT,
		ALTER COLUMN idtipo TYPE TEXT,
		ALTER COLUMN tipo TYPE TEXT,
		ALTER COLUMN nomecodigo TYPE TEXT,
		ALTER COLUMN "Cabecalho" TYPE TEXT,
		ALTER COLUMN titulo TYPE TEXT,
		ALTER COLUMN titulotexto TYPE TEXT,
		ALTER COLUMN titulo_label TYPE TEXT,
		ALTER COLUMN capitulo TYPE TEXT,
		ALTER COLUMN capitulotexto TYPE TEXT,
		ALTER COLUMN capitulo_label TYPE TEXT,
		ALTER COLUMN secao TYPE TEXT,
		ALTER COLUMN secaotexto TYPE TEXT,
		ALTER COLUMN secao_label TYPE TEXT,
		ALTER COLUMN subsecao TYPE TEXT,
		ALTER COLUMN subsecaotexto TYPE TEXT,
		ALTER COLUMN subsecao_label TYPE TEXT,
		ALTER COLUMN num_artigo TYPE TEXT,
		ALTER COLUMN "Artigos" TYPE TEXT`).Error; err != nil {
		return err
	}

	if err := db.Exec(`ALTER TABLE vade_mecum_jurisprudencia
		ALTER COLUMN id TYPE TEXT,
		ALTER COLUMN idtipo TYPE TEXT,
		ALTER COLUMN tipo TYPE TEXT,
		ALTER COLUMN idcodigo TYPE TEXT,
		ALTER COLUMN nomecodigo TYPE TEXT,
		ALTER COLUMN "Cabecalho" TYPE TEXT,
		ALTER COLUMN "Tipo" TYPE TEXT,
		ALTER COLUMN idramo TYPE TEXT,
		ALTER COLUMN ramotexto TYPE TEXT,
		ALTER COLUMN idassunto TYPE TEXT,
		ALTER COLUMN assuntotexto TYPE TEXT,
		ALTER COLUMN idenunciado TYPE TEXT,
		ALTER COLUMN "Enunciado" TYPE TEXT,
		ALTER COLUMN idsecao TYPE TEXT,
		ALTER COLUMN secao TYPE TEXT,
		ALTER COLUMN secaotexto TYPE TEXT,
		ALTER COLUMN idsubsecao TYPE TEXT,
		ALTER COLUMN subsecao TYPE TEXT,
		ALTER COLUMN subsecaotexto TYPE TEXT,
		ALTER COLUMN num_artigo TYPE TEXT,
		ALTER COLUMN "Normativo" TYPE TEXT,
		ALTER COLUMN "Ordem" TYPE TEXT`).Error; err != nil {
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

	if !migrator.HasConstraint(&model.VadeMecumCodigo{}, "vade_mecum_codigos_idcodigo_key") {
		if err := db.Exec(`ALTER TABLE vade_mecum_codigos ADD CONSTRAINT vade_mecum_codigos_idcodigo_key UNIQUE (idcodigo)`).Error; err != nil {
			if !strings.Contains(err.Error(), "already exists") {
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
