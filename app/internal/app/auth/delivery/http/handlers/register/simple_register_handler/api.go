package simple_register_handler

import (
	"net/http"

	"github.com/sirupsen/logrus"

	user_usecase "getme-backend/internal/app/auth/usecase"
	"getme-backend/internal/pkg/handler/handler_errors"
	"getme-backend/internal/pkg/utilits/delivery"
	postgresql_utilits "getme-backend/internal/pkg/utilits/postgresql"
)

var codeByError = delivery.CodeMap{
	user_usecase.SimpleAuthExists: {
		http.StatusConflict, handler_errors.UserAlreadyExists, logrus.InfoLevel},
	user_usecase.NicknameExists: {
		http.StatusConflict, handler_errors.UserAlreadyExists, logrus.InfoLevel},
	user_usecase.LoginExists: {
		http.StatusConflict, handler_errors.UserAlreadyExists, logrus.InfoLevel},
	user_usecase.UserExist: {
		http.StatusConflict, handler_errors.UserAlreadyExists, logrus.InfoLevel},
	user_usecase.IncorrectLoginOrPassword: {
		http.StatusUnprocessableEntity, handler_errors.IncorrectLoginOrPassword, logrus.InfoLevel},
	postgresql_utilits.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
}
