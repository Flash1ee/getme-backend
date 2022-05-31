package plans_task_handler

import (
	"net/http"

	"github.com/sirupsen/logrus"

	plans_usecase "getme-backend/internal/app/plans/usecase"
	"getme-backend/internal/pkg/handler/handler_errors"
	"getme-backend/internal/pkg/utilits/delivery"
	postgresql_utilits "getme-backend/internal/pkg/utilits/postgresql"
)

var codeByErrGET = delivery.CodeMap{
	plans_usecase.UnknownRole: {
		http.StatusBadRequest, handler_errors.UnknownRole, logrus.WarnLevel},
	postgresql_utilits.NotFound: {
		http.StatusNotFound, handler_errors.UserNotFound, logrus.WarnLevel},
	postgresql_utilits.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
}
