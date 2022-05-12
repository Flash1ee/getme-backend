package token_handler

import (
	"net/http"

	"github.com/sirupsen/logrus"

	token_redis_repository "getme-backend/internal/app/token/repository/redis"
	"getme-backend/internal/pkg/handler/handler_errors"
	"getme-backend/internal/pkg/utilits/delivery"
	"getme-backend/internal/pkg/utilits/postgresql"
)

var codeByErrorGET = delivery.CodeMap{
	token_redis_repository.SetError: {
		http.StatusInternalServerError, handler_errors.InternalError, logrus.ErrorLevel},

	postgresql_utilits.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.InternalError, logrus.ErrorLevel},
}
