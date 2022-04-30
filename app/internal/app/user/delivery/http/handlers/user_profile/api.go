package user_profile_handler

import (
	"getme-backend/internal/app/user/repository"
	"getme-backend/internal/pkg/handler/handler_errors"
	"getme-backend/internal/pkg/utilits/delivery"
	"getme-backend/internal/pkg/utilits/postgresql"
	"github.com/sirupsen/logrus"
	"net/http"
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
	repository.NicknameAlreadyExist: {
		http.StatusConflict, repository.NicknameAlreadyExist, logrus.WarnLevel},
	repository.EmailAlreadyExist: {
		http.StatusConflict, repository.EmailAlreadyExist, logrus.WarnLevel},
}
