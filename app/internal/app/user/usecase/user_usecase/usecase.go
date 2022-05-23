package user_usecase

import (
	"getme-backend/internal/app/user/dto"
	"getme-backend/internal/app/user/entities"
	"getme-backend/internal/app/user/repository"
	user_usecase "getme-backend/internal/app/user/usecase"
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

func (u *UserUsecase) FindByNickname(nickname string) (*entities.User, error) {
	res, err := u.userRepository.FindByNickname(nickname)
	if err == postgresql_utilits.NotFound {
		return nil, user_usecase.UserNotFound
	}
	return res, err
}

func (u *UserUsecase) CreateBaseUser(nickname string) (int64, error) {
	return u.userRepository.CreateBaseUser(nickname)
}

func (u *UserUsecase) CreateFilledUser(data *dto.UserUsecase) (int64, error) {
	us, err := u.userRepository.CreateFilledUser(data.ToUserEntity())

	return us, err
}
func (u *UserUsecase) FindByID(id int64) (*dto.UserUsecase, error) {
	user, err := u.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return dto.ToUserUsecase(user), nil
}
