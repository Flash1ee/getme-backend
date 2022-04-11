package usecase

import "getme-backend/internal/app/user/dto"

type Usecase interface {
	Create(us *dto.UserRequest) ([]dto.UserResponse, error)
	Get(nickname string) (*dto.UserResponse, error)
	Update(us *dto.UserRequest) (*dto.UserResponse, error)
}
