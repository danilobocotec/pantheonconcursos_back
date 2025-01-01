package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	docs "github.com/thepantheon/api/docs"
	"github.com/thepantheon/api/internal/config"
	"github.com/thepantheon/api/internal/handler"
	"github.com/thepantheon/api/pkg/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// Override Swagger host/scheme based on environment configuration
	swaggerHost := cfg.Server.Host
	if swaggerHost == "" {
		swaggerHost = fmt.Sprintf("localhost:%s", cfg.Server.Port)
	}
	docs.SwaggerInfo.Host = swaggerHost
	if cfg.Server.Scheme != "" {
		docs.SwaggerInfo.Schemes = []string{cfg.Server.Scheme}
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

		cursos := api.Group("/cursos")
		{
			cursos.GET("", handlers.GetCourses)
			cursos.POST("", handlers.CreateCourse)
			cursos.GET("/categorias", handlers.GetCourseCategories)
			cursos.POST("/categorias", handlers.CreateCourseCategory)
			cursos.PUT("/categorias/:id", handlers.UpdateCourseCategory)
			cursos.DELETE("/categorias/:id", handlers.DeleteCourseCategory)
			cursos.PUT("/:id", handlers.UpdateCourse)
			cursos.DELETE("/:id", handlers.DeleteCourse)
		}

		meusCursos := api.Group("/meus-cursos")
		{
			meusCursos.GET("/modulos", handlers.GetMyModules)
			meusCursos.POST("/modulos", handlers.CreateCourseModuleStandalone)
			meusCursos.GET("/itens", handlers.GetMyItems)
			meusCursos.POST("/itens", handlers.CreateCourseItemStandalone)
			meusCursos.PUT("/itens/:id", handlers.UpdateCourseItem)
			meusCursos.DELETE("/itens/:id", handlers.DeleteCourseItem)
			meusCursos.PUT("/modulos/:id", handlers.UpdateCourseModule)
			meusCursos.DELETE("/modulos/:id", handlers.DeleteCourseModule)
		}

		questoes := api.Group("/questoes")
		{
			questoes.GET("", handlers.GetQuestoes)
			questoes.GET("/filtros", handlers.GetQuestaoFilters)
			questoes.POST("", handlers.CreateQuestao)
			questoes.GET("/:id", handlers.GetQuestaoByID)
			questoes.PUT("/:id", handlers.UpdateQuestao)
			questoes.DELETE("/:id", handlers.DeleteQuestao)
		}

		meuDesempenho := api.Group("/meu-desempenho")
		{
			meuDesempenho.GET("", handlers.GetUserPerformance)
			meuDesempenho.GET("/resumo", handlers.GetUserPerformanceSummary)
			meuDesempenho.POST("", handlers.CreateUserPerformance)
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
			codigos.POST("/import/estatuto", handlers.ImportEstatuto)
			codigos.GET("/capas", handlers.GetCapasVadeMecumCodigo)
			codigos.POST("/capas", handlers.CreateCapaVadeMecumCodigo)
			codigos.PUT("/capas/:id", handlers.UpdateCapaVadeMecumCodigo)
			codigos.GET("/grouped", handlers.GetCodigosGrouped)
			codigos.GET("/:id", handlers.GetCodigoByID)
			codigos.PUT("/:id", handlers.UpdateCodigo)
			codigos.DELETE("/:id", handlers.DeleteCodigo)
		}

		estatutos := api.Group("/vade-mecum/estatutos")
		{
			estatutos.GET("", handlers.GetEstatutos)
			estatutos.GET("/gruposervico", handlers.GetEstatutoGrupoServico)
			estatutos.POST("", handlers.CreateEstatuto)
			estatutos.GET("/:id", handlers.GetEstatutoByID)
			estatutos.PUT("/:id", handlers.UpdateEstatuto)
			estatutos.DELETE("/:id", handlers.DeleteEstatuto)
		}

		constituicao := api.Group("/vade-mecum/constituicao")
		{
			constituicao.GET("", handlers.GetConstituicoes)
			constituicao.GET("/gruposervico", handlers.GetConstituicaoGrupoServico)
			constituicao.POST("", handlers.CreateConstituicao)
			constituicao.GET("/:id", handlers.GetConstituicaoByID)
			constituicao.PUT("/:id", handlers.UpdateConstituicao)
			constituicao.DELETE("/:id", handlers.DeleteConstituicao)
			constituicao.POST("/import", handlers.ImportConstituicao)
		}

		leis := api.Group("/vade-mecum/leis")
		{
			leis.GET("", handlers.GetLeis)
			leis.GET("/gruposervico", handlers.GetLeiGrupoServico)
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
			oab.PUT("/capas/:id", handlers.UpdateCapaVadeMecumOAB)
			oab.POST("/import", handlers.ImportVadeMecumOAB)
		}

		juris := api.Group("/vade-mecum/jurisprudencia")
		{
			juris.GET("", handlers.GetVadeMecumJurisprudencia)
			juris.GET("/grouped", handlers.GetVadeMecumJurisprudenciaGrouped)
			juris.POST("", handlers.CreateVadeMecumJurisprudencia)
			juris.GET("/capas", handlers.GetCapasVadeMecumJurisprudencia)
			juris.POST("/capas", handlers.CreateCapaVadeMecumJurisprudencia)
			juris.PUT("/capas/:id", handlers.UpdateCapaVadeMecumJurisprudencia)
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
