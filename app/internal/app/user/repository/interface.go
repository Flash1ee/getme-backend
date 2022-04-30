package user_repository

import (
	"context"

	"getme-backend/internal/app/user/entities"
)

//go:generate mockgen -destination=mock/$GOFILE -package=mock -source=$GOFILE

type Repository interface {
	Create(ctx context.Context, user *entities.User) (*entities.User, error)
	GetUserByTelegramID(ctx context.Context, tgID int64) (*entities.User, error)
}
