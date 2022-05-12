package token_usecase

import (
	"getme-backend/internal/app/token/dto"
)

//go:generate mockgen -destination=mock/$GOFILE -package=mock -source=$GOFILE

type Usecase interface {
	// GetToken with Errors:
	//		app.GeneralError with Errors
	//			token_redis_repository.SetError
	GetToken(userID int64) (dto.TokenResponse, error)
	// GetTokenByData with Errors:
	//		app.GeneralError with Errors
	//			token_redis_repository.SetError
	GetTokenByData(tokenSources dto.TokenSourcesUsecase) (dto.TokenResponse, error)
	//	CheckToken with Errors:
	//	token_redis_repository.NotFound
	//	app.GeneralError with Errors
	//		token_redis_repository.InvalidStorageData
	CheckToken(token dto.TokenUsecase) (bool, error)
	//	CheckTokenByUser with Errors:
	//		InvalidUserToken
	//		token_redis_repository.NotFound
	//		app.GeneralError with Errors
	//			token_redis_repository.InvalidStorageData
	CheckTokenByUser(token dto.TokenUsecase, userID int64) error
	Check(identifierData dto.TokenSourcesUsecase, token dto.TokenUsecase) error
}
