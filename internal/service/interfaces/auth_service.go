package interfaces

import "zurihaqi.github.io-backend/internal/model"

type AuthService interface {
	Register(name, email, password string) (*model.User, error)
	Login(email, password string) (string, error)
}
