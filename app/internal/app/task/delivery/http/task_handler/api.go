package task_handler

import (
	"net/http"

	"github.com/sirupsen/logrus"

	task_usecase "getme-backend/internal/app/task/usecase"
	"getme-backend/internal/pkg/handler/handler_errors"
	"getme-backend/internal/pkg/utilits/delivery"
	postgresql_utilits "getme-backend/internal/pkg/utilits/postgresql"
)

var codeByErrPOST = delivery.CodeMap{
	task_usecase.UserHaveNotThisPlan: {
		http.StatusBadRequest, handler_errors.MentorHaveNotThisPlan, logrus.WarnLevel},
	postgresql_utilits.NotFound: {
		http.StatusNotFound, handler_errors.UserNotFound, logrus.WarnLevel},
	postgresql_utilits.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
}
