package token_usecase

import (
	"strconv"
	"time"

	"getme-backend/internal/app/token/dto"
	"getme-backend/internal/app/token/entities"
	"getme-backend/internal/app/token/repository"
	token_usecase "getme-backend/internal/app/token/usecase"

	uuid "github.com/satori/go.uuid"
)

const (
	timeExp        = time.Hour * 3
	expiredJWTTime = time.Minute * 15
)

type TokenUsecase struct {
	repository    token_repository.Repository
	repositoryJWT token_repository.RepositoryJWT
}

func NewTokenUsecase(repository token_repository.Repository, repoJWT token_repository.RepositoryJWT) *TokenUsecase {
	return &TokenUsecase{
		repository:    repository,
		repositoryJWT: repoJWT,
	}
}

// GetToken with Errors:
//		app.GeneralError with Errors
//			token_redis_repository.SetError
func (u *TokenUsecase) GetToken(userID int64) (dto.TokenResponse, error) {
	token := uuid.NewV4().String()
	userIDtoStr := strconv.Itoa(int(userID))
	err := u.repository.Set(token, userIDtoStr, int(timeExp.Seconds()))

	if err != nil {
		return dto.TokenResponse{}, err
	}
	tokenModel := &entities.Token{Token: token}
	return *dto.ToTokenResponse(tokenModel), nil
}

// GetTokenByData with Errors:
//		use JWT for auth token generating
//		app.GeneralError with Errors
// 			token_jwt_repository.ErrorSignedToken
func (u *TokenUsecase) GetTokenByData(tokenSources dto.TokenSourcesUsecase) (dto.TokenResponse, error) {
	convertedData := tokenSources.ToTokenSourcesEntity(expiredJWTTime)
	res, err := u.repositoryJWT.Create(*convertedData)
	if err != nil {
		return dto.TokenResponse{}, err
	}
	return *dto.ToTokenResponse(&res), nil
}

// Check Errors:
//      token_jwt_repository.BadToken
// 		app.GeneralError with Error
// 			token_jwt_repository.ParseClaimsError
//			token_jwt_repository.TokenExpired
func (u *TokenUsecase) Check(identifierData dto.TokenSourcesUsecase, token dto.TokenUsecase) error {
	sources := identifierData.ToTokenSourcesEntity(expiredJWTTime)
	return u.repositoryJWT.Check(*sources, *token.ToTokenEntity())

}

//	CheckToken with Errors:
//	token_redis_repository.NotFound
//	app.GeneralError with Errors
//		token_redis_repository.InvalidStorageData
func (u *TokenUsecase) CheckToken(token dto.TokenUsecase) (bool, error) {
	_, err := u.repository.Get(token.Token)
	if err != nil {
		return false, err
	}
	return true, nil
}

//	CheckTokenByUser with Errors:
//		token_redis_repository.NotFound
//		token_usecase.InvalidUserToken
//		app.GeneralError with Errors
//			token_redis_repository.InvalidStorageData
func (u *TokenUsecase) CheckTokenByUser(token dto.TokenUsecase, userID int64) error {
	userTokenID, err := u.repository.Get(token.Token)
	if err != nil {
		return err
	}
	userTokenIDToInt, err := strconv.Atoi(userTokenID)
	if err != nil {
		return err
	}
	if int64(userTokenIDToInt) != userID {
		return token_usecase.InvalidUserToken
	}

	return nil
}
