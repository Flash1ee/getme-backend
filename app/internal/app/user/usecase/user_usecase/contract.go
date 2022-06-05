package user_usecase

import "getme-backend/internal/app/auth/dto"

//go:generate mockgen -destination=mock/$GOFILE -package=mock -source=$GOFILE

type authChecker interface {
	Check(data *dto.AuthUsecase) bool
}
