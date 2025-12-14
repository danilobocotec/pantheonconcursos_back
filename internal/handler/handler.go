package handler

import (
	"github.com/thepantheon/api/internal/repository"
	"github.com/thepantheon/api/internal/service"
	"gorm.io/gorm"
)

type Handlers struct {
	userService       *service.UserService
	authService       *service.AuthService
	socialAuthService *service.SocialAuthService
	planService       *service.PlanService
	adminSecret       string
	vadeMecumService  *service.VadeMecumService
	codigoService     *service.VadeMecumCodigoService
}

func NewHandlers(db *gorm.DB, googleClientID, googleClientSecret, facebookAppID, facebookAppSecret, redirectURL, jwtSecret, adminSecret string) *Handlers {
	userRepo := repository.NewUserRepository(db)
	planRepo := repository.NewPlanRepository(db)
	vadeMecumRepo := repository.NewVadeMecumRepository(db)
	codigoRepo := repository.NewVadeMecumCodigoRepository(db)
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userService, jwtSecret)
	socialAuthService := service.NewSocialAuthService(userService, googleClientID, googleClientSecret, facebookAppID, facebookAppSecret, redirectURL)
	planService := service.NewPlanService(planRepo)
	vadeMecumService := service.NewVadeMecumService(vadeMecumRepo)
	codigoService := service.NewVadeMecumCodigoService(codigoRepo)

	return &Handlers{
		userService:       userService,
		authService:       authService,
		socialAuthService: socialAuthService,
		planService:       planService,
		vadeMecumService:  vadeMecumService,
		codigoService:     codigoService,
		adminSecret:       adminSecret,
	}
}
