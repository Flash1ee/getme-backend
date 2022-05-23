package telegram_auth_handler

import (
	"net/http"

	"github.com/sirupsen/logrus"

	token_jwt_repository "getme-backend/internal/app/token/repository/jwt"
	"getme-backend/internal/pkg/handler/handler_errors"
	"getme-backend/internal/pkg/utilits/delivery"
	"getme-backend/internal/pkg/utilits/postgresql"
)

var codesByErrorsGET = delivery.CodeMap{
	token_jwt_repository.BadToken: {
		http.StatusBadRequest, handler_errors.TokenInvalid, logrus.WarnLevel,
	},
	token_jwt_repository.ParseClaimsError: {
		http.StatusBadRequest, handler_errors.TokenInvalid, logrus.WarnLevel,
	},
	token_jwt_repository.TokenExpired: {
		http.StatusBadRequest, handler_errors.TokenInvalid, logrus.WarnLevel},
	postgresql_utilits.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
	postgresql_utilits.NotFound: {
		http.StatusNotFound, handler_errors.UserNotFound, logrus.WarnLevel},
}

var codesByErrorsPOST = delivery.CodeMap{
	postgresql_utilits.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
	postgresql_utilits.NotFound: {
		http.StatusNotFound, handler_errors.UserNotFound, logrus.WarnLevel},
}
