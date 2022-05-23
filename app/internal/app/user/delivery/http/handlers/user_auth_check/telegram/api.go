package user_telegram_auth

import (
	"net/http"

	"github.com/sirupsen/logrus"

	repository_postgresql "getme-backend/internal/app/user/repository/postgresql"
	user_usecase "getme-backend/internal/app/user/usecase"
	"getme-backend/internal/pkg/handler/handler_errors"
	"getme-backend/internal/pkg/utilits/delivery"
	"getme-backend/internal/pkg/utilits/postgresql"
)

var codesByErrorsGET = delivery.CodeMap{
	user_usecase.ArgError: {
		http.StatusInternalServerError, handler_errors.InternalError, logrus.InfoLevel,
	},
	user_usecase.BadAuth: {
		http.StatusBadRequest, handler_errors.InvalidQueries, logrus.InfoLevel,
	},
	repository_postgresql.CreateError: {
		http.StatusServiceUnavailable, handler_errors.InternalError, logrus.ErrorLevel,
	},
	postgresql_utilits.DefaultErrDB: {
		http.StatusServiceUnavailable, handler_errors.BDError, logrus.ErrorLevel},
}
