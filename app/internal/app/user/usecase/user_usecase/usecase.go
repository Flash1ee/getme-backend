package user_usecase

import (
	"context"

	"getme-backend/internal/app/user/dto"
	"getme-backend/internal/app/user/repository"
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
func (u *UserUsecase) Auth(user *dto.UserAuthUsecase) (*dto.UserAuthUsecase, error) {
	if user == nil {
		return nil, ArgError
	}
	ok := u.authChecker.Check(user)
	if !ok {
		return nil, BadAuth
	}
	userFromDB, err := u.userRepository.Create(context.Background(), user.ToUserEntity())
	if err != nil {
		return nil, err
	}

	return dto.ToUserAuthUsecase(userFromDB), nil
}
