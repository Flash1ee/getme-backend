package user_usecase

import (
	"getme-backend/internal/app/user/dto"
)

type Usecase interface {
	Auth(user *dto.UserAuthUsecase) (*dto.UserResponse, error)
}
