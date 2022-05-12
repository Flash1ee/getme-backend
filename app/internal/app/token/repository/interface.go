package token_repository

import "getme-backend/internal/app/token/entities"

//go:generate mockgen -destination=mock/$GOFILE -package=mock -source=$GOFILE

type Repository interface {
	// Set Errors:
	// 		app.GeneralError with Errors
	// 			repository_redis.SetError
	Set(key string, value string, timeExp int) error
	// Get Errors:
	//		repository_redis.NotFound
	// 		app.GeneralError with Errors
	// 			repository_redis.InvalidStorageData
	Get(key string) (string, error)
}

type RepositoryJWT interface {
	// Check Errors:
	// 		BadToken
	// 		app.GeneralError with Error
	// 			ParseClaimsError
	// 			TokenExpired
	Check(sources entities.TokenSources, tokenString entities.Token) error

	// Create Errors:
	// 		app.GeneralError with Error
	// 			ErrorSignedToken
	Create(sources entities.TokenSources) (entities.Token, error)
}
