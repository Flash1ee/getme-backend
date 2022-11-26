package jwt_auth_handler

import (
	"net/http"

	"github.com/sirupsen/logrus"

	auth_usecase "getme-backend/internal/app/auth/usecase"
	"getme-backend/internal/pkg/handler/handler_errors"
	"getme-backend/internal/pkg/utilits/delivery"
	postgresql_utilits "getme-backend/internal/pkg/utilits/postgresql"
)

var codesByErrors = delivery.CodeMap{
	postgresql_utilits.NotFound: {
		http.StatusNotFound, handler_errors.UserNotFound, logrus.WarnLevel},
	postgresql_utilits.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
	auth_usecase.IncorrectLoginOrPassword: {
		http.StatusUnauthorized, handler_errors.IncorrectLoginOrPassword, logrus.InfoLevel},
}
