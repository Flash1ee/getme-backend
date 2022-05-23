package user_usecase

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"getme-backend/internal/app"
	"getme-backend/internal/app/user/dto"
	"getme-backend/internal/app/user/repository"
	userUsecase "getme-backend/internal/app/user/usecase"
	"getme-backend/internal/pkg/handler/handler_errors"
	"getme-backend/internal/pkg/usecase"
	postgresql_utilits "getme-backend/internal/pkg/utilits/postgresql"
)

type UserUsecase struct {
	usecase.BaseUsecase
	userRepository user_repository.Repository
	authChecker    authChecker
}

func NewUserUsecase(repo user_repository.Repository, authCheck authChecker) *UserUsecase {
	return &UserUsecase{
		userRepository: repo,
		authChecker:    authCheck,
	}
}

// AuthTelegram Errors:
//		user_usecase.ArgError
//		user_usecase.BadAuth
// 		app.GeneralError with Error
// 			repository_postgresql.CreateError
func (u *UserUsecase) AuthTelegram(user *dto.UserAuthUsecase) (*dto.UserAuthUsecase, error) {
	if user == nil {
		return nil, userUsecase.ArgError
	}
	ok := u.authChecker.Check(user)
	if !ok {
		return nil, userUsecase.BadAuth
	}
	userFromDB, err := u.userRepository.Create(context.Background(), user.ToUserEntity())
	if err != nil {
		return nil, errors.Wrap(err, "UserUsecase: AuthTelegram(): create user error")
	}

	return dto.ToUserAuthUsecase(userFromDB), nil
}

func (u *UserUsecase) AuthSimple(login string, password string) (int64, error) {
	user, err := u.userRepository.FindByLoginSimple(login)
	if err != nil {
		return -1, err
	}

	usecaseUser := dto.ToUserRegisterUsecase(user)

	if !usecaseUser.ComparePassword(password) {
		return -1, handler_errors.IncorrectLoginOrPassword
	}
	return user.UserID, nil
}

// CreateSimple Errors:
//		models.EmptyPassword
// 		models.IncorrectEmailOrPassword
//		user_repository.LoginAlreadyExist
//		user_usecase.UserExist
// 		app.GeneralError with Errors
// 			repository.DefaultErrDB
func (u *UserUsecase) CreateSimple(user *dto.UserSimpleRegistrationUsecase) (int64, error) {
	checkUser, err := u.userRepository.FindByLoginSimple(user.Login)
	if err != nil && err != postgresql_utilits.NotFound {
		return -1, errors.Wrap(err, fmt.Sprintf("error on create user with login %v", user.Login))
	}

	if checkUser != nil {
		return -1, userUsecase.UserExist
	}

	if err = user.Encrypt(); err != nil {
		if errors.Is(err, userUsecase.EmptyPassword) {
			return -1, err
		}

		return -1, app.GeneralError{
			Err:         userUsecase.BadEncrypt,
			ExternalErr: err,
		}
	}
	userFromDB, err := u.userRepository.CreateSimple(user.ToUserRegisterEntity())
	if err != nil {
		return -1, err
	}

	return userFromDB.ID, nil
}
