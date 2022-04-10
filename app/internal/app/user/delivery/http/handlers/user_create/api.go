package user_create_handler

import (
	"getme-backend/internal/pkg/handler/handler_errors"
	"getme-backend/internal/pkg/utilits/delivery"
	"getme-backend/internal/pkg/utilits/postgresql"
	"github.com/sirupsen/logrus"
	"net/http"
)

var codesByErrorsPOST = delivery.CodeMap{
	postgresql_utilits.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
}
