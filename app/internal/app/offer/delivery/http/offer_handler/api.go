package offer_handler

import (
	"net/http"

	"github.com/sirupsen/logrus"

	offer_usecase "getme-backend/internal/app/offer/usecase"
	skill_usecase "getme-backend/internal/app/skill/usecase"
	"getme-backend/internal/pkg/handler/handler_errors"
	"getme-backend/internal/pkg/utilits/delivery"
	postgresql_utilits "getme-backend/internal/pkg/utilits/postgresql"
)

var codeByErrGET = delivery.CodeMap{
	offer_usecase.NotMentor: {
		http.StatusNotFound, UserNotMentor, logrus.ErrorLevel},
	postgresql_utilits.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
}
var codeByErrPOST = delivery.CodeMap{
	offer_usecase.LogicError: {
		http.StatusBadRequest, LogicError, logrus.ErrorLevel},
	offer_usecase.AlreadyExists: {
		http.StatusConflict, OfferAlreadyExists, logrus.ErrorLevel},
	skill_usecase.SkillNotExists: {
		http.StatusNotFound, SkillNotFound, logrus.ErrorLevel},
	offer_usecase.MentorNotExist: {
		http.StatusNotFound, MentorNotFound, logrus.ErrorLevel},
	postgresql_utilits.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
}
