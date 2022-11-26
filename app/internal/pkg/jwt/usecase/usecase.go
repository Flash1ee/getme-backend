package usecase_csrf

import (
	"getme-backend/internal/pkg/jwt/models"
)

type Usecase interface {
	Check(userId int64, token string) error
	Create(userId int64) (models.Token, error)
}
