package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/thepantheon/api/internal/config"
	"github.com/thepantheon/api/internal/handler"
	"github.com/thepantheon/api/pkg/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/thepantheon/api/docs"
)

// @title           ThePantheon API
// @version         1.0
// @description     API REST para sistema de concursos com autenticação JWT
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.email  support@pantheonconcursos.com.br

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Digite "Bearer" seguido do token JWT

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
	handlers := handler.NewHandlers(
		db,
		cfg.OAuth.GoogleClientID,
		cfg.OAuth.GoogleClientSecret,
		cfg.OAuth.FacebookAppID,
		cfg.OAuth.FacebookAppSecret,
		cfg.OAuth.RedirectURL,
		cfg.JWT.Secret,
	)

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
			// Traditional auth
			auth.POST("/login", handlers.Login)
			auth.POST("/register", handlers.Register)
			auth.POST("/refresh", middleware.AuthMiddleware(), handlers.RefreshToken)

			// Social auth
			auth.POST("/social", handlers.SocialLogin)
			auth.GET("/google/url", handlers.GoogleAuthURL)
			auth.GET("/google/callback", handlers.GoogleCallback)
			auth.GET("/facebook/url", handlers.FacebookAuthURL)
			auth.GET("/facebook/callback", handlers.FacebookCallback)
		}

		plans := api.Group("/plans")
		{
			plans.GET("", handlers.GetPlans)
		}
	}

	// Start server
	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	log.Printf("Server starting on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
