package handler

import (
	"github.com/thepantheon/api/internal/repository"
	"github.com/thepantheon/api/internal/service"
	"gorm.io/gorm"
)

type Handlers struct {
	userService           *service.UserService
	authService           *service.AuthService
	socialAuthService     *service.SocialAuthService
	planService           *service.PlanService
	adminSecret           string
	vadeMecumService      *service.VadeMecumService
	codigoService         *service.VadeMecumCodigoService
	leisService           *service.VadeMecumLeiService
	capaCodigoService     *service.CapaVadeMecumCodigoService
	oabService            *service.VadeMecumOABService
	capaOABService        *service.CapaVadeMecumOABService
	jurisprudenciaService *service.VadeMecumJurisprudenciaService
	capaJurisService      *service.CapaVadeMecumJurisprudenciaService
}

func NewHandlers(db *gorm.DB, googleClientID, googleClientSecret, facebookAppID, facebookAppSecret, redirectURL, jwtSecret, adminSecret string) *Handlers {
	userRepo := repository.NewUserRepository(db)
	planRepo := repository.NewPlanRepository(db)
	vadeMecumRepo := repository.NewVadeMecumRepository(db)
	codigoRepo := repository.NewVadeMecumCodigoRepository(db)
	leisRepo := repository.NewVadeMecumLeiRepository(db)
	capaCodigoRepo := repository.NewCapaVadeMecumCodigoRepository(db)
	oabRepo := repository.NewVadeMecumOABRepository(db)
	capaOABRepo := repository.NewCapaVadeMecumOABRepository(db)
	jurisRepo := repository.NewVadeMecumJurisprudenciaRepository(db)
	capaJurisRepo := repository.NewCapaVadeMecumJurisprudenciaRepository(db)
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userService, jwtSecret)
	socialAuthService := service.NewSocialAuthService(userService, googleClientID, googleClientSecret, facebookAppID, facebookAppSecret, redirectURL)
	planService := service.NewPlanService(planRepo)
	vadeMecumService := service.NewVadeMecumService(vadeMecumRepo)
	codigoService := service.NewVadeMecumCodigoService(codigoRepo)
	leisService := service.NewVadeMecumLeiService(leisRepo)
	capaCodigoService := service.NewCapaVadeMecumCodigoService(capaCodigoRepo)
	oabService := service.NewVadeMecumOABService(oabRepo)
	capaOABService := service.NewCapaVadeMecumOABService(capaOABRepo)
	jurisprudenciaService := service.NewVadeMecumJurisprudenciaService(jurisRepo)
	capaJurisService := service.NewCapaVadeMecumJurisprudenciaService(capaJurisRepo)

	return &Handlers{
		userService:           userService,
		authService:           authService,
		socialAuthService:     socialAuthService,
		planService:           planService,
		vadeMecumService:      vadeMecumService,
		codigoService:         codigoService,
		leisService:           leisService,
		adminSecret:           adminSecret,
		capaCodigoService:     capaCodigoService,
		oabService:            oabService,
		capaOABService:        capaOABService,
		jurisprudenciaService: jurisprudenciaService,
		capaJurisService:      capaJurisService,
	}
}
