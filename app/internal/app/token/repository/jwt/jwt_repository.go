package token_jwt_repository

import (
	"time"

	"getme-backend/internal/app"
	"getme-backend/internal/app/token/entities"

	uuid "github.com/satori/go.uuid"

	"github.com/golang-jwt/jwt"
)

var (
	secretToken = uuid.NewV4()
)

type JwtRepository struct {
	Secret []byte
}

type jwtCsrfClaims struct {
	jwt.StandardClaims
	IdentifierData string
}

func NewJwtRepository() *JwtRepository {
	return &JwtRepository{Secret: secretToken.Bytes()}
}

func (tk *JwtRepository) parseClaims(token *jwt.Token) (interface{}, error) {
	method, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok || method.Alg() != "HS256" {
		return nil, IncorrectTokenSigningMethod
	}
	return tk.Secret, nil
}

// Check Errors:
// 		repository_jwt.BadToken
// 		app.GeneralError with Error
// 			repository_jwt.ParseClaimsError
// 			repository_jwt.TokenExpired
func (tk *JwtRepository) Check(sources entities.TokenSources, token entities.Token) error {
	claims := &jwtCsrfClaims{}
	tokenParsing, err := jwt.ParseWithClaims(token.Token, claims, tk.parseClaims)
	if err != nil || !tokenParsing.Valid {
		retErr := &app.GeneralError{ExternalErr: err}

		switch err.(*jwt.ValidationError).Errors {
		case jwt.ValidationErrorExpired:
			retErr.Err = TokenExpired
		case jwt.ValidationErrorUnverifiable:
			retErr.ExternalErr = IncorrectTokenSigningMethod
			retErr.Err = ParseClaimsError
		default:
			retErr.Err = ParseClaimsError
		}
		return retErr
	}

	if claims.IdentifierData != sources.IdentifierData {
		return BadToken
	}
	return nil
}

// Create Errors:
// 		app.GeneralError with Error
// 			repository_jwt.ErrorSignedToken
func (tk *JwtRepository) Create(sources entities.TokenSources) (entities.Token, error) {
	data := jwtCsrfClaims{
		IdentifierData: sources.IdentifierData,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: sources.ExpiredTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	res, err := token.SignedString(tk.Secret)
	if err != nil {
		return entities.Token{}, &app.GeneralError{
			Err:         ErrorSignedToken,
			ExternalErr: err,
		}
	}
	return entities.Token{Token: res}, nil
}
