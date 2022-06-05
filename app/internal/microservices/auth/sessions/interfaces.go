package sessions

import "getme-backend/internal/microservices/auth/sessions/models"

//go:generate mockgen -destination=mock/$GOFILE -package=mock -source=$GOFILE

type SessionRepository interface {
	Set(session *models.Session) error
	GetUserId(key string, updExpiration int) (string, error)
	Del(session *models.Session) error
}

type SessionUsecase interface {
	Check(uniqID string) (models.Result, error)
	Create(userID int64) (models.Result, error)
	Delete(uniqID string) error
	// CreateByTokenID gen session by user token
	CreateByTokenID(tokenID string, userID int64) (models.ResultByToken, error)
	CheckWithDelete(tokenID string) (models.ResultByToken, error)
}
