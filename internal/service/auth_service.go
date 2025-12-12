package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/thepantheon/api/internal/model"
)

type AuthService struct {
	userService *UserService
	jwtSecret   string
}

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func NewAuthService(userService *UserService, jwtSecret string) *AuthService {
	return &AuthService{
		userService: userService,
		jwtSecret:   jwtSecret,
	}
}

func (s *AuthService) Login(req *model.LoginRequest) (*model.LoginResponse, error) {
	user, err := s.userService.GetUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := s.userService.VerifyPassword(user, req.Password); err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, expiresIn, err := s.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		Token:     token,
		ExpiresIn: expiresIn,
		User:      *user,
	}, nil
}

func (s *AuthService) GenerateToken(user *model.User) (string, int64, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: user.ID.String(),
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", 0, err
	}

	return tokenString, expirationTime.Unix(), nil
}

func (s *AuthService) ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
