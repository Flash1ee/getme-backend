package dto

import (
	"golang.org/x/crypto/bcrypt"

	"getme-backend/internal/app/user/entities"
)

type UserAuthUsecase struct {
	ID         int64  `json:"id"`
	TelegramID int64  `json:"tg_id"`
	AuthDate   int64  `json:"auth_date"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Username   string `json:"username"`
	Avatar     string `json:"photo_url"`
	Hash       string `json:"hash"`
}

func (req *UserAuthUsecase) ToUserEntity() *entities.User {
	return &entities.User{
		TelegramID: req.TelegramID,
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		Nickname:   req.Username,
		Avatar:     req.Avatar,
	}
}

func ToUserAuthUsecase(user *entities.User) *UserAuthUsecase {
	return &UserAuthUsecase{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Nickname,
		Avatar:    user.Avatar,
	}
}

type UserSimpleRegistrationUsecase struct {
	Login             string `json:"login"`
	Password          string `json:"password"`
	EncryptedPassword string
}

func (req *UserSimpleRegistrationUsecase) ToUserRegisterEntity() *entities.UserSimpleAuth {
	return &entities.UserSimpleAuth{
		Login:             req.Login,
		EncryptedPassword: req.EncryptedPassword,
	}
}
func ToUserRegisterUsecase(user *entities.UserSimpleAuth) *UserSimpleRegistrationUsecase {
	return &UserSimpleRegistrationUsecase{
		Login:             user.Login,
		EncryptedPassword: user.EncryptedPassword,
	}
}

func (u *UserSimpleRegistrationUsecase) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil
}

func (u *UserSimpleRegistrationUsecase) encryptString(s string) (string, error) {
	enc, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(enc), nil
}
func (u *UserSimpleRegistrationUsecase) Encrypt() error {
	enc, err := u.encryptString(u.Password)
	if err != nil {
		return err
	}
	u.EncryptedPassword = enc
	return nil
}
