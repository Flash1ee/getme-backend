package usecase_csrf

import (
	"time"

	"getme-backend/internal/pkg/jwt/models"
	repository_token "getme-backend/internal/pkg/jwt/repository"
)

var expiredJWTTime = time.Minute * 15

type JWTUsecase struct {
	repository repository_token.Repository
}

func NewJWTUsecase(repo repository_token.Repository) *JWTUsecase {
	return &JWTUsecase{
		repository: repo,
	}
}

func (u *JWTUsecase) Check(userId int64, token string) error {
	sources := models.TokenSources{
		UserId: userId,
	}
	return u.repository.Check(sources, models.Token(token))

}

// Create Errors:
// 		app.GeneralError with Error
// 			repository_jwt.ErrorSignedToken
func (u *JWTUsecase) Create(userId int64) (models.Token, error) {
	data := models.TokenSources{
		UserId:      userId,
		ExpiredTime: time.Now().Add(expiredJWTTime),
	}
	return u.repository.Create(data)
}
