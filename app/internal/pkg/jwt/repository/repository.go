package repository_token

import "getme-backend/internal/pkg/jwt/models"

type Repository interface {
	Check(sources models.TokenSources, tokenString models.Token) error
	Create(sources models.TokenSources) (models.Token, error)
}
