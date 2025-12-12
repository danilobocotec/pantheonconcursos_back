package main

import (
	"fmt"
	"log"

	"github.com/thepantheon/api/internal/config"
	"github.com/thepantheon/api/internal/handler"
	"github.com/thepantheon/api/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database
	db, err := config.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Auto migrate models
	config.AutoMigrate(db)

	// Create Gin router
	router := gin.Default()

	// Apply middlewares
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.ErrorHandlingMiddleware())
	router.Use(middleware.LoggingMiddleware())

	// Initialize handlers
	handlers := handler.NewHandlers(db)

	// API Routes
	api := router.Group("/api/v1")
	{
		// Health check
		api.GET("/health", handlers.HealthCheck)

		// Users routes
		users := api.Group("/users")
		{
			users.POST("", handlers.CreateUser)
			users.GET("", handlers.GetUsers)
			users.GET("/:id", handlers.GetUser)
			users.PUT("/:id", handlers.UpdateUser)
			users.DELETE("/:id", handlers.DeleteUser)
		}

		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/login", handlers.Login)
			auth.POST("/register", handlers.Register)
			auth.POST("/refresh", middleware.AuthMiddleware(), handlers.RefreshToken)
		}
	}

	// Start server
	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	log.Printf("Server starting on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
