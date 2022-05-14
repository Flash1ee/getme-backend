package user_usecase

import (
	"context"

	"github.com/pkg/errors"

	"getme-backend/internal/app/user/dto"
	"getme-backend/internal/app/user/repository"
	userUsecase "getme-backend/internal/app/user/usecase"
	"getme-backend/internal/pkg/usecase"
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

// Auth Errors:
//		user_usecase.ArgError
//		user_usecase.BadAuth
// 		app.GeneralError with Error
// 			repository_postgresql.CreateError
func (u *UserUsecase) Auth(user *dto.UserAuthUsecase) (*dto.UserAuthUsecase, error) {
	if user == nil {
		return nil, userUsecase.ArgError
	}
	ok := u.authChecker.Check(user)
	if !ok {
		return nil, userUsecase.BadAuth
	}
	userFromDB, err := u.userRepository.Create(context.Background(), user.ToUserEntity())
	if err != nil {
		return nil, errors.Wrap(err, "UserUsecase: Auth(): create user error")
	}

	return dto.ToUserAuthUsecase(userFromDB), nil
}
