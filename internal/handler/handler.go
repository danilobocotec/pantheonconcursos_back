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
	questaoService        *service.QuestaoService
	userPerformanceService *service.UserPerformanceService
	courseService         *service.CourseService
	vadeMecumService      *service.VadeMecumService
	codigoService         *service.VadeMecumCodigoService
	leisService           *service.VadeMecumLeiService
	capaCodigoService     *service.CapaVadeMecumCodigoService
	oabService            *service.VadeMecumOABService
	capaOABService        *service.CapaVadeMecumOABService
	jurisprudenciaService *service.VadeMecumJurisprudenciaService
	capaJurisService      *service.CapaVadeMecumJurisprudenciaService
	estatutoService       *service.VadeMecumEstatutoService
	constituicaoService   *service.VadeMecumConstituicaoService
}

func NewHandlers(db *gorm.DB, googleClientID, googleClientSecret, facebookAppID, facebookAppSecret, redirectURL, jwtSecret, adminSecret string) *Handlers {
	userRepo := repository.NewUserRepository(db)
	planRepo := repository.NewPlanRepository(db)
	questaoRepo := repository.NewQuestaoRepository(db)
	userPerformanceRepo := repository.NewUserPerformanceRepository(db)
	courseRepo := repository.NewCourseRepository(db)
	vadeMecumRepo := repository.NewVadeMecumRepository(db)
	codigoRepo := repository.NewVadeMecumCodigoRepository(db)
	estatutoRepo := repository.NewVadeMecumEstatutoRepository(db)
	constituicaoRepo := repository.NewVadeMecumConstituicaoRepository(db)
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
	questaoService := service.NewQuestaoService(questaoRepo)
	userPerformanceService := service.NewUserPerformanceService(userPerformanceRepo)
	courseService := service.NewCourseService(courseRepo)
	vadeMecumService := service.NewVadeMecumService(vadeMecumRepo)
	codigoService := service.NewVadeMecumCodigoService(codigoRepo)
	estatutoService := service.NewVadeMecumEstatutoService(estatutoRepo)
	constituicaoService := service.NewVadeMecumConstituicaoService(constituicaoRepo)
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
		questaoService:        questaoService,
		userPerformanceService: userPerformanceService,
		courseService:         courseService,
		vadeMecumService:      vadeMecumService,
		codigoService:         codigoService,
		leisService:           leisService,
		adminSecret:           adminSecret,
		capaCodigoService:     capaCodigoService,
		oabService:            oabService,
		capaOABService:        capaOABService,
		jurisprudenciaService: jurisprudenciaService,
		capaJurisService:      capaJurisService,
		estatutoService:       estatutoService,
		constituicaoService:   constituicaoService,
	}
}
