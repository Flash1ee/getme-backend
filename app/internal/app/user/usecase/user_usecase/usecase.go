package usecase

import (
	"getme-backend/internal/app/user/dto"
	"getme-backend/internal/app/user/repository"
)

type UserUsecase struct {
	userRepository repository.Repository
}

func (u *UserUsecase) Create(us *dto.UserRequest) ([]dto.UserResponse, error) {
	return []dto.UserResponse{}, nil
}

func (u *UserUsecase) Get(nickname string) (*dto.UserResponse, error) {
	return nil, nil
}

func (u *UserUsecase) Update(us *dto.UserRequest) (*dto.UserResponse, error) {
	return nil, nil
}
