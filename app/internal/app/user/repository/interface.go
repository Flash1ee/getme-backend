package user_repository

import (
	"getme-backend/internal/app/user/entities"
)

type Repository interface {
	Create(us *entities.User) ([]entities.User, error)
	Update(us *entities.User) (*entities.User, error)
	Get(nickname string) (*entities.User, error)
	Delete(nickname string) error
}
