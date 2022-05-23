package user_profile_handler

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"getme-backend/internal/pkg/handler/handler_errors"
	"getme-backend/internal/pkg/utilits/delivery"
	postgresql_utilits "getme-backend/internal/pkg/utilits/postgresql"
)

var codesByErrorsGET = delivery.CodeMap{
	postgresql_utilits.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
}
