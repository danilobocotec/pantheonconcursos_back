package repository

import (
	"github.com/google/uuid"
	"github.com/thepantheon/api/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetByID(id uuid.UUID) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetAll(limit, offset int) ([]model.User, error) {
	var users []model.User
	err := r.db.Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}

func (r *UserRepository) Update(id uuid.UUID, user *model.User) error {
	return r.db.Model(&model.User{}, "id = ?", id).Updates(user).Error
}

func (r *UserRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&model.User{}, "id = ?", id).Error
}

func (r *UserRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&model.User{}).Count(&count).Error
	return count, err
}

func (r *UserRepository) Exists(email string) (bool, error) {
	var count int64
	err := r.db.Model(&model.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}
