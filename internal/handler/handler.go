package handler

import (
	"github.com/thepantheon/api/internal/repository"
	"github.com/thepantheon/api/internal/service"
	"gorm.io/gorm"
)

type Handlers struct {
	userService *service.UserService
	authService *service.AuthService
}

func NewHandlers(db *gorm.DB) *Handlers {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userService, "your_jwt_secret")

	return &Handlers{
		userService: userService,
		authService: authService,
	}
}
