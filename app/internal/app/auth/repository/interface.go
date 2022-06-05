package auth_repository

import (
	"getme-backend/internal/app/auth/entities"
)

//go:generate mockgen -destination=mock/$GOFILE -package=mock -source=$GOFILE

type Repository interface {
	// FindByLoginSimple Errors:
	// 		postgresql_utilits.NotFound
	// 		app.GeneralError with Errors:
	// 			postgresql_utilits.DefaultErrDB
	FindByLoginSimple(login string) (*entities.SimpleAuth, error)

	// CreateSimple Errors:
	// 		EmailAlreadyExist
	// 		NicknameAlreadyExist
	// 		app.GeneralError with Errors
	// 			postgresql_utilits.DefaultErrDB
	CreateSimple(auth *entities.SimpleAuth) (*entities.SimpleAuth, error)

	UpdateTelegramAuthTime(tgID int64) error
	// FindByTelegramID Errors:
	//		postgresql_utilits.NotFound
	// 		app.GeneralError with Errors:
	// 			postgresql_utilits.DefaultErrDB
	FindByTelegramID(tgID int64) (*entities.TelegramAuth, error)
	CreateTelegramAuthRecord(auth *entities.TelegramAuth) (*entities.TelegramAuth, error)

	//// Create Errors:
	//// 		app.GeneralError with Error
	//// 			repository_postgresql.CreateError
	//Create(ctx context.Context, user *entities.User) (*entities.User, error)
	//// CreateWithUpdate Errors:
	//// 		app.GeneralError with Error
	//// 			repository_postgresql.CreateError
	//CreateWithUpdate(ctx context.Context, user *entities.User) (*entities.User, error)
	//// GetUserByTelegramID Errors:
	//// 		app.GeneralError with Error
	//// 			repository_postgresql.GetError
	//GetUserByTelegramID(ctx context.Context, tgID int64) (*entities.User, error)
	//// CheckExistsByTelegram Errors:
	//// 		app.GeneralError with Errors:
	//// 			postgresql_utilits.DefaultErrDB
	//CheckExistsByTelegram(tgID int64) (bool, error)
	//

}
