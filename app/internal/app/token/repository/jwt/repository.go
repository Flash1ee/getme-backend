package token_jwt_repository

import "getme-backend/internal/app/token/entities"

//go:generate mockgen -destination=mock/$GOFILE -package=mock -source=$GOFILE

type Repository interface {
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
