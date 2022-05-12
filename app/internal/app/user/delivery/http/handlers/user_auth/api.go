package user_auth_handler

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"getme-backend/internal/pkg/handler/handler_errors"
	"getme-backend/internal/pkg/utilits/delivery"
	"getme-backend/internal/pkg/utilits/postgresql"
)

var codesByErrorsGET = delivery.CodeMap{
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
