package token_jwt_repository

import "errors"

var (
	ParseClaimsError            = errors.New("can not parse body of token")
	ErrorSignedToken            = errors.New("can not sign token by secret key")
	IncorrectTokenSigningMethod = errors.New("incorrect parsing token signing method")
	TokenExpired                = errors.New("token expired")
	BadToken                    = errors.New("data in token are invalid")
)
