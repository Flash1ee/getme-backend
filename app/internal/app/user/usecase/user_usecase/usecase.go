package user_usecase

import (
	"getme-backend/internal/app/task/repository"
	"getme-backend/internal/app/user/dto"
)

type UserUsecase struct {
	userRepository repository.Repository
	authChecker    authChecker
}

func NewUserUsecase(repo repository.Repository, authCheck authChecker) *UserUsecase {
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
	return &dto.UserResponse{
		Nickname: user.Username,
		Fullname: user.FirstName + " " + user.LastName,
	}, nil
}
