package offer_id_accept_handler

import (
	"net/http"

	"github.com/sirupsen/logrus"

	offer_usecase "getme-backend/internal/app/offer/usecase"
	"getme-backend/internal/pkg/handler/handler_errors"
	"getme-backend/internal/pkg/utilits/delivery"
	postgresql_utilits "getme-backend/internal/pkg/utilits/postgresql"
)

var codeByErrPOST = delivery.CodeMap{
	offer_usecase.InvalidOfferID: {
		http.StatusBadRequest, handler_errors.LogicError, logrus.ErrorLevel},
	postgresql_utilits.NotFound: {
		http.StatusNotFound, handler_errors.OfferNotFound, logrus.WarnLevel},
	postgresql_utilits.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
}

var codeByErrDELETE = delivery.CodeMap{
	offer_usecase.InvalidOfferID: {
		http.StatusBadRequest, handler_errors.LogicError, logrus.ErrorLevel},
	postgresql_utilits.NotFound: {
		http.StatusNotFound, handler_errors.OfferNotFound, logrus.WarnLevel},
	postgresql_utilits.DefaultErrDB: {
		http.StatusInternalServerError, handler_errors.BDError, logrus.ErrorLevel},
}
