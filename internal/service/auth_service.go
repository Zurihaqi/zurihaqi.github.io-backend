package service

import (
	"errors"

	"zurihaqi.github.io-backend/internal/model"
	repoInterface "zurihaqi.github.io-backend/internal/repository/interfaces"
	"zurihaqi.github.io-backend/internal/service/interfaces"
	"zurihaqi.github.io-backend/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo repoInterface.UserRepository
}

func NewAuthService(userRepo repoInterface.UserRepository) interfaces.AuthService {
	return &AuthService{UserRepo: userRepo}
}

func (s *AuthService) Register(name, email, password string) (*model.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Name:     name,
		Email:    email,
		Password: string(hashed),
	}

	return s.UserRepo.Create(user)
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
