package repository

import (
	"gorm.io/gorm"
	"zurihaqi.github.io-backend/internal/model"
	repoInterface "zurihaqi.github.io-backend/internal/repository/interfaces"
)

type userRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) repoInterface.UserRepository {
	return &userRepositoryImpl{DB: db}
}

func (r *userRepositoryImpl) Create(user *model.User) (*model.User, error) {
	err := r.DB.Create(user).Error
	return user, err
}

func (r *userRepositoryImpl) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
