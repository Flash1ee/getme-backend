package usecase

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"getme-backend/internal/microservices/auth/sessions"
	"getme-backend/internal/microservices/auth/sessions/models"

	"golang.org/x/crypto/sha3"

	uuid "github.com/satori/go.uuid"
)

const (
	ExpiredCookiesTime = 48 * time.Hour
	UnknownUser        = -1
	//
	ExpiredTemporaryCookiesTime = 1 * time.Minute
	UnknownToken                = "invalid token"
)

type SessionUsecase struct {
	sessionRepository sessions.SessionRepository
}

func NewSessionUsecase(sessionRep sessions.SessionRepository) *SessionUsecase {
	return &SessionUsecase{
		sessionRepository: sessionRep,
	}
}

func (u *SessionUsecase) Create(userID int64) (models.Result, error) {
	strUserID := fmt.Sprintf("%d", userID)

	session := &models.Session{
		UserID:     strUserID,
		UniqID:     generateUniqID(strUserID),
		Expiration: int(ExpiredCookiesTime.Milliseconds()),
	}
	if err := u.sessionRepository.Set(session); err != nil {
		return models.Result{UserID: UnknownUser}, err
	}
	return models.Result{UserID: userID, UniqID: session.UniqID}, nil
}

// CreateByTokenID - Видоизмененный Create: UserID и UniqID поменяны местами.
// Используется только в качестве временной сессии.
// Для создания сессий использовать Create!
func (u *SessionUsecase) CreateByTokenID(tokenID string, userID int64) (models.ResultByToken, error) {
	strUserID := fmt.Sprintf("%d", userID)
	session := &models.Session{
		UserID:     strUserID,
		UniqID:     tokenID,
		Expiration: int(ExpiredTemporaryCookiesTime.Milliseconds()),
	}
	if err := u.sessionRepository.Set(session); err != nil {
		return models.ResultByToken{TokenID: string(rune(UnknownUser))}, err
	}

	return models.ResultByToken{TokenID: session.UniqID, UserID: session.UserID}, nil
}

func (u *SessionUsecase) Delete(uniqID string) error {
	session := &models.Session{
		UniqID: uniqID,
	}
	return u.sessionRepository.Del(session)
}

func (u *SessionUsecase) Check(uniqID string) (models.Result, error) {
	userID, err := u.sessionRepository.GetUserId(uniqID, int(ExpiredCookiesTime.Milliseconds()))
	if err != nil {
		return models.Result{UserID: UnknownUser, UniqID: uniqID}, err
	}

	intUserID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return models.Result{UserID: UnknownUser, UniqID: uniqID}, err
	}
	return models.Result{UserID: intUserID, UniqID: uniqID}, nil
}
func (u *SessionUsecase) CheckWithDelete(tokenID string) (models.ResultByToken, error) {
	userID, err := u.sessionRepository.GetUserId(tokenID, 0)
	if err != nil {
		return models.ResultByToken{TokenID: UnknownToken, UserID: string(rune(UnknownUser))}, err
	}

	return models.ResultByToken{TokenID: tokenID, UserID: userID}, nil
}

func generateUniqID(userID string) string {
	value := append([]byte(userID), uuid.NewV4().Bytes()...)
	hash := sha3.Sum512(value)

	return hex.EncodeToString(hash[:])
}
