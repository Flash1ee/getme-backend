package dto

import (
	"golang.org/x/crypto/bcrypt"

	"getme-backend/internal/app/auth/entities"
)

type AuthUsecase struct {
	ID         int64  `json:"id"`
	TelegramID int64  `json:"tg_id"`
	AuthDate   int64  `json:"auth_date"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Username   string `json:"username"`
	Avatar     string `json:"photo_url"`
	Hash       string `json:"hash"`
}

func (req *AuthUsecase) ToTelegramEntity() *entities.TelegramAuth {
	return &entities.TelegramAuth{
		TelegramID: req.TelegramID,
		UserID:     req.ID,
	}
}

func ToTelegramAuthUsecase(tg *entities.TelegramAuth) *AuthUsecase {
	return &AuthUsecase{
		ID:         tg.UserID,
		TelegramID: tg.TelegramID,
	}
}

type SimpleRegistrationUsecase struct {
	Login             string `json:"login"`
	Password          string `json:"password"`
	EncryptedPassword string
	UserID            int64
}

func (req *SimpleRegistrationUsecase) ToUserRegisterEntity() *entities.SimpleAuth {
	return &entities.SimpleAuth{
		Login:             req.Login,
		EncryptedPassword: req.EncryptedPassword,
		UserID:            req.UserID,
	}
}
func ToUserRegisterUsecase(user *entities.SimpleAuth) *SimpleRegistrationUsecase {
	return &SimpleRegistrationUsecase{
		Login:             user.Login,
		EncryptedPassword: user.EncryptedPassword,
	}
}

func (u *SimpleRegistrationUsecase) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil
}

func (u *SimpleRegistrationUsecase) encryptString(s string) (string, error) {
	enc, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(enc), nil
}
func (u *SimpleRegistrationUsecase) Encrypt() error {
	enc, err := u.encryptString(u.Password)
	if err != nil {
		return err
	}
	u.EncryptedPassword = enc
	return nil
}
