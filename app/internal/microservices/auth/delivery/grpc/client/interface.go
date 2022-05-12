package client

import (
	"context"

	"getme-backend/internal/microservices/auth/sessions/models"
)

//go:generate mockgen -destination=mock/$GOFILE -package=mock -source=$GOFILE

type AuthCheckerClient interface {
	Check(ctx context.Context, sessionID string) (models.Result, error)
	Create(ctx context.Context, userID int64) (models.Result, error)
	Delete(ctx context.Context, sessionID string) error
	// CreateByToken create session by string token
	CreateByToken(ctx context.Context, tokenID string, userID int64) (models.ResultByToken, error)
	// CheckWithDelete check token session by string token
	CheckWithDelete(ctx context.Context, tokenID string) (models.ResultByToken, error)
}
