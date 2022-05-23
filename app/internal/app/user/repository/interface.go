package user_repository

import (
	"context"

	"getme-backend/internal/app/user/entities"
)

//go:generate mockgen -destination=mock/$GOFILE -package=mock -source=$GOFILE

type Repository interface {
	// Create Errors:
	// 		app.GeneralError with Error
	// 			repository_postgresql.CreateError
	Create(ctx context.Context, user *entities.User) (*entities.User, error)
	// CreateWithUpdate Errors:
	// 		app.GeneralError with Error
	// 			repository_postgresql.CreateError
	CreateWithUpdate(ctx context.Context, user *entities.User) (*entities.User, error)
	// GetUserByTelegramID Errors:
	// 		app.GeneralError with Error
	// 			repository_postgresql.GetError
	GetUserByTelegramID(ctx context.Context, tgID int64) (*entities.User, error)
	// CheckExists Errors:
	// 		app.GeneralError with Errors:
	// 			postgresql_utilits.DefaultErrDB
	CheckExists(tgID int64) (bool, error)
	// FindByLoginSimple Errors:
	// 		postgresql_utilits.NotFound
	// 		app.GeneralError with Errors:
	// 			postgresql_utilits.DefaultErrDB
	FindByLoginSimple(string) (*entities.UserSimpleAuth, error)
	// CreateSimple Errors:
	// 		LoginAlreadyExist
	// 		NicknameAlreadyExist
	// 		app.GeneralError with Errors
	// 			repository.DefaultErrDB
	CreateSimple(auth *entities.UserSimpleAuth) (*entities.UserSimpleAuth, error)
}
