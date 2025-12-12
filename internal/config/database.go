package config

import (
	"log"

	"github.com/thepantheon/api/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *Config) (*gorm.DB, error) {
	dsn := cfg.Database.GetDSN()
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Database connection established successfully")
	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		// Add more models here as needed
	)
}
