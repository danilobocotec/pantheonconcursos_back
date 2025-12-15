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
		cfg.Admin.Secret,
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
			auth.POST("/admin/register", handlers.RegisterAdmin)
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

		vade := api.Group("/vade-mecum")
		{
			vade.GET("", handlers.GetVadeMecum)
			vade.POST("", handlers.CreateVadeMecum)
			vade.GET("/:id", handlers.GetVadeMecumByID)
			vade.PUT("/:id", handlers.UpdateVadeMecum)
			vade.DELETE("/:id", handlers.DeleteVadeMecum)
		}

		vadeCategory := api.Group("/vade-mecum/category/:category")
		{
			vadeCategory.GET("", handlers.GetVadeMecumByCategory)
			vadeCategory.POST("", handlers.CreateVadeMecumByCategory)
			vadeCategory.PUT("/:id", handlers.UpdateVadeMecumByCategory)
			vadeCategory.DELETE("/:id", handlers.DeleteVadeMecumByCategory)
		}

		codigos := api.Group("/vade-mecum/codigos")
		{
			codigos.GET("", handlers.GetCodigos)
			codigos.POST("", handlers.CreateCodigo)
			codigos.POST("/import", handlers.ImportCodigos)
			codigos.GET("/capas", handlers.GetCapasVadeMecumCodigo)
			codigos.POST("/capas", handlers.CreateCapaVadeMecumCodigo)
			codigos.PUT("/capas/:nomecodigo", handlers.UpdateCapaVadeMecumCodigo)
			codigos.GET("/grouped", handlers.GetCodigosGrouped)
			codigos.GET("/:id", handlers.GetCodigoByID)
			codigos.PUT("/:id", handlers.UpdateCodigo)
			codigos.DELETE("/:id", handlers.DeleteCodigo)
		}

		leis := api.Group("/vade-mecum/leis")
		{
			leis.GET("", handlers.GetLeis)
			leis.POST("", handlers.CreateLei)
			leis.GET("/:id", handlers.GetLeiByID)
			leis.PUT("/:id", handlers.UpdateLei)
			leis.DELETE("/:id", handlers.DeleteLei)
			leis.POST("/import", handlers.ImportLeis)
		}

		oab := api.Group("/vade-mecum/oab")
		{
			oab.GET("", handlers.GetVadeMecumOAB)
			oab.POST("", handlers.CreateVadeMecumOAB)
			oab.GET("/:id", handlers.GetVadeMecumOABByID)
			oab.PUT("/:id", handlers.UpdateVadeMecumOAB)
			oab.DELETE("/:id", handlers.DeleteVadeMecumOAB)
			oab.GET("/capas", handlers.GetCapasVadeMecumOAB)
			oab.POST("/capas", handlers.CreateCapaVadeMecumOAB)
			oab.PUT("/capas/:nomecodigo", handlers.UpdateCapaVadeMecumOAB)
			oab.POST("/import", handlers.ImportVadeMecumOAB)
		}

		juris := api.Group("/vade-mecum/jurisprudencia")
		{
			juris.GET("", handlers.GetVadeMecumJurisprudencia)
			juris.POST("", handlers.CreateVadeMecumJurisprudencia)
			juris.GET("/capas", handlers.GetCapasVadeMecumJurisprudencia)
			juris.POST("/capas", handlers.CreateCapaVadeMecumJurisprudencia)
			juris.PUT("/capas/:nomecodigo", handlers.UpdateCapaVadeMecumJurisprudencia)
			juris.POST("/import", handlers.ImportVadeMecumJurisprudencia)
			juris.GET("/:id", handlers.GetVadeMecumJurisprudenciaByID)
			juris.PUT("/:id", handlers.UpdateVadeMecumJurisprudencia)
			juris.DELETE("/:id", handlers.DeleteVadeMecumJurisprudencia)
		}
	}

	// Start server
	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	log.Printf("Server starting on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
