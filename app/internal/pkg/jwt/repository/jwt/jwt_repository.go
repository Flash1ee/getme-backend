package repository_jwt

import (
	"time"

	"getme-backend/internal/app"
	"getme-backend/internal/pkg/jwt/models"

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
	UserId int64 `json:"user_id"`
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
func (tk *JwtRepository) Check(sources models.TokenSources, tokenString models.Token) error {
	claims := &jwtCsrfClaims{}
	token, err := jwt.ParseWithClaims(string(tokenString), claims, tk.parseClaims)
	if err != nil || !token.Valid {
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

	if claims.UserId != sources.UserId {
		return BadToken
	}
	return nil
}

// Create Errors:
// 		app.GeneralError with Error
// 			repository_jwt.ErrorSignedToken
func (tk *JwtRepository) Create(sources models.TokenSources) (models.Token, error) {
	data := jwtCsrfClaims{
		UserId: sources.UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: sources.ExpiredTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	res, err := token.SignedString(tk.Secret)
	if err != nil {
		return "", &app.GeneralError{
			Err:         ErrorSignedToken,
			ExternalErr: err,
		}
	}
	return models.Token(res), nil
}
