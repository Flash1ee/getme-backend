package user_simple_auth

import (
	"net/http"

	"github.com/sirupsen/logrus"

	user_usecase "getme-backend/internal/app/user/usecase"
	"getme-backend/internal/pkg/handler/handler_errors"
	"getme-backend/internal/pkg/utilits/delivery"
	postgresql_utilits "getme-backend/internal/pkg/utilits/postgresql"
)

var codesByErrors = delivery.CodeMap{
	postgresql_utilits.NotFound: {
		http.StatusNotFound, handler_errors.UserNotFound, logrus.WarnLevel},
	postgresql_utilits.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
	user_usecase.IncorrectEmailOrPassword: {
		http.StatusUnauthorized, handler_errors.IncorrectLoginOrPassword, logrus.InfoLevel},
}
