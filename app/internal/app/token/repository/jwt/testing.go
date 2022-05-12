package token_jwt_repository

import (
	"testing"
	"time"

	"getme-backend/internal/app/token/entities"
)

func TestSources(t *testing.T) *entities.TokenSources {
	t.Helper()
	return &entities.TokenSources{
		IdentifierData: "127.0.0.1",
		ExpiredTime:    time.Now().Add(time.Minute),
	}
}
