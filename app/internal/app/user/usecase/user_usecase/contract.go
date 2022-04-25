package user_usecase

import "getme-backend/internal/app/user/dto"

type authChecker interface {
	Check(data *dto.UserAuthUsecase) bool
}
