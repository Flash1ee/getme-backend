package auth_usecase

import (
	"fmt"

	"github.com/pkg/errors"

	"getme-backend/internal/app"
	"getme-backend/internal/app/auth/dto"
	auth_repository "getme-backend/internal/app/auth/repository"
	auth_usecase "getme-backend/internal/app/auth/usecase"
	user_repository "getme-backend/internal/app/user/repository"
	"getme-backend/internal/pkg/usecase"
	postgresql_utilits "getme-backend/internal/pkg/utilits/postgresql"
)

type AuthUsecase struct {
	usecase.BaseUsecase
	authRepository auth_repository.Repository
	userRepository user_repository.Repository
	authChecker    authChecker
}

func NewAuthUsecase(repoAuth auth_repository.Repository, repoUser user_repository.Repository, authCheck authChecker) *AuthUsecase {
	return &AuthUsecase{
		authRepository: repoAuth,
		userRepository: repoUser,
		authChecker:    authCheck,
	}
}

//AuthSimple with Errors:
//	auth_usecase.IncorrectLoginOrPassword
func (u *AuthUsecase) AuthSimple(login string, password string) (int64, error) {
	userAuth, err := u.authRepository.FindByLoginSimple(login)
	if err != nil {
		return -1, errors.Wrapf(err, "AuthUsecase: AuthSimple = FindByLogin error, login = %v", login)
	}

	usecaseUser := dto.ToUserRegisterUsecase(userAuth)

	if !usecaseUser.ComparePassword(password) {
		return -1, auth_usecase.IncorrectLoginOrPassword
	}
	return userAuth.UserID, nil
}

// AuthTelegram Errors:
//		auth_usecase.ArgError
//		auth_usecase.BadAuth
// 		app.GeneralError with Error
// 			repository_postgresql.CreateError
//			postgresql_utilits.DefaultErrDB
func (u *AuthUsecase) AuthTelegram(user *dto.AuthUsecase) (*dto.AuthUsecase, error) {
	if user == nil {
		return nil, auth_usecase.ArgError
	}
	ok := u.authChecker.Check(user)
	if !ok {
		return nil, auth_usecase.BadAuth
	}
	// Check exists user!

	if _, err := u.authRepository.FindByTelegramID(user.TelegramID); err != nil {
		if err == postgresql_utilits.NotFound {
			res, err := u.authRepository.CreateTelegramAuthRecord(user.ToTelegramEntity())
			if err != nil {
				return nil, errors.Wrapf(err,
					"AuthUsecase: AuthTelegram(): CreateTelegramAuthRecord with tg_id = %v failed", user.TelegramID)
			}
			return dto.ToTelegramAuthUsecase(res), nil
		}
	}

	err := u.authRepository.UpdateTelegramAuthTime(user.TelegramID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// CreateSimple Errors:
//		auth_usecase.EmptyPassword
// 		auth_usecase.IncorrectEmailOrPassword
//		user_repository.LoginAlreadyExist
//		auth_usecase.SimpleAuthExists
// 		app.GeneralError with Errors
// 			repository.DefaultErrDB
func (u *AuthUsecase) CreateSimple(user *dto.SimpleRegistrationUsecase) (int64, error) {
	checkUser, err := u.authRepository.FindByLoginSimple(user.Login)
	if err != nil && err != postgresql_utilits.NotFound {
		return -1, errors.Wrap(err, fmt.Sprintf("error on create user with login %v", user.Login))
	}

	if checkUser != nil {
		return -1, auth_usecase.SimpleAuthExists
	}

	if err = user.Encrypt(); err != nil {
		if errors.Is(err, auth_usecase.EmptyPassword) {
			return -1, err
		}

		return -1, app.GeneralError{
			Err:         auth_usecase.BadEncrypt,
			ExternalErr: err,
		}
	}

	userAuthSimpleData, err := u.authRepository.CreateSimple(user.ToUserRegisterEntity())
	if err != nil {
		return -1, err
	}

	return userAuthSimpleData.ID, nil
}
