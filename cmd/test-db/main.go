package main

import (
	"fmt"
	"log"

	"github.com/thepantheon/api/internal/config"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	fmt.Println("Connecting to Supabase PostgreSQL...")
	fmt.Printf("Host: %s\n", cfg.Database.Host)
	fmt.Printf("Port: %s\n", cfg.Database.Port)
	fmt.Printf("Database: %s\n", cfg.Database.DBName)
	fmt.Printf("User: %s\n", cfg.Database.User)
	fmt.Printf("SSL Mode: %s\n", cfg.Database.SSLMode)

	// Initialize database
	db, err := config.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("✅ Successfully connected to Supabase PostgreSQL!")

	// Auto migrate models
	if err := config.AutoMigrate(db); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("✅ Database migration completed!")

	// Test query
	var result string
	if err := db.Raw("SELECT NOW()").Scan(&result).Error; err != nil {
		log.Fatalf("Failed to execute test query: %v", err)
	}

	fmt.Printf("\n✅ Test query result: %s\n", result)
}
