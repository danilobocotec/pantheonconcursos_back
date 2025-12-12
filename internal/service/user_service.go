package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/thepantheon/api/internal/model"
	"github.com/thepantheon/api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(req *model.CreateUserRequest) (*model.User, error) {
	// Confirm password validation is enforced by binding tag eqfield=Password
	// Check if user already exists
	exists, err := s.repo.Exists(req.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:       uuid.New(),
		Email:    req.Email,
		FullName: req.FullName,
		Password: string(hashedPassword),
		Active:   true,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUserByID(id uuid.UUID) (*model.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) GetUserByEmail(email string) (*model.User, error) {
	return s.repo.GetByEmail(email)
}

func (s *UserService) GetAllUsers(limit, offset int) ([]model.User, error) {
	return s.repo.GetAll(limit, offset)
}

func (s *UserService) UpdateUser(id uuid.UUID, req *model.UpdateUserRequest) (*model.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Email != "" {
		user.Email = req.Email
	}
	if req.FullName != "" {
		user.FullName = req.FullName
	}
	user.Active = req.Active

	if err := s.repo.Update(id, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) DeleteUser(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *UserService) VerifyPassword(user *model.User, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
