package user_usecase

import (
	"context"

	"getme-backend/internal/app/user/dto"
	"getme-backend/internal/app/user/repository"
)

type UserUsecase struct {
	userRepository user_repository.Repository
	authChecker    authChecker
}

func NewUserUsecase(repo user_repository.Repository, authCheck authChecker) *UserUsecase {
	return &UserUsecase{
		userRepository: repo,
		authChecker:    authCheck,
	}
}
func (u *UserUsecase) Auth(user *dto.UserAuthUsecase) (*dto.UserResponse, error) {
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

	return &dto.UserResponse{
		Nickname: userFromDB.Nickname,
		Fullname: userFromDB.FirstName + " " + userFromDB.LastName,
	}, nil
}
