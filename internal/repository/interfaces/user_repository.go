package interfaces

import "zurihaqi.github.io-backend/internal/model"

type UserRepository interface {
	Create(user *model.User) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}
