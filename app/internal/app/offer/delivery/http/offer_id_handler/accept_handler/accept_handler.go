package offer_id_accept_handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"getme-backend/internal/app/offer/dto"
	offer_usecase "getme-backend/internal/app/offer/usecase"
	plan_dto "getme-backend/internal/app/plan/dto"
	session_client "getme-backend/internal/microservices/auth/delivery/grpc/client"
	middleware2 "getme-backend/internal/microservices/auth/sessions/middleware"
	"getme-backend/internal/pkg/adapter/echo_adapter"
	bh "getme-backend/internal/pkg/handler"
	"getme-backend/internal/pkg/handler/handler_errors"
)

type AcceptHandler struct {
	sessionClient session_client.AuthCheckerClient
	offerUsecase  offer_usecase.Usecase
	bh.BaseHandler
}

func NewAcceptHandler(log *logrus.Logger, sClient session_client.AuthCheckerClient, offerUs offer_usecase.Usecase) *AcceptHandler {
	h := &AcceptHandler{
		BaseHandler:   *bh.NewBaseHandler(log),
		sessionClient: sClient,
		offerUsecase:  offerUs,
	}

	h.AddMethod(http.MethodPost, h.POST, echo_adapter.WrapMiddlewareToFunc(middleware2.NewSessionMiddleware(sClient, log).CheckFunc))
	h.AddMethod(http.MethodDelete, h.DELETE, echo_adapter.WrapMiddlewareToFunc(middleware2.NewSessionMiddleware(sClient, log).CheckFunc))

	return h
}

func (h *AcceptHandler) POST(ctx echo.Context) error {
	userID, ok := ctx.Request().Context().Value("user_id").(int64)
	if !ok {
		h.Log(ctx.Request()).Error("can not get user_id from context")
		h.Error(ctx, http.StatusInternalServerError, handler_errors.InternalError)
		return handler_errors.InternalError
	}
	req := &dto.RequestAcceptOffer{}

	err := h.GetRequestBody(ctx, req)
	if err != nil {
		h.Log(ctx.Request()).Warnf("can not parse request %s", err)
		h.Error(ctx, http.StatusUnprocessableEntity, handler_errors.InvalidBody)
		return handler_errors.InvalidBody
	}
	offerID, _, err := h.GetInt64FromParam(ctx, "id")
	if err != nil {
		h.Log(ctx.Request()).Warnf("OfferIDAcceptHandler: can not parse path param ID %s", err)
		h.Error(ctx, http.StatusBadRequest, handler_errors.InvalidParameters)
		return handler_errors.InvalidParameters
	}

	usecaseDTO := dto.ToOfferAcceptUsecaseDTO(*req)
	usecaseDTO.OfferID = offerID

	res, err := h.offerUsecase.Accept(userID, usecaseDTO)
	if err != nil {
		h.Log(ctx.Request()).Warnf("OfferIDAcceptHandler: can not parse path param ID %s", err)
		h.UsecaseError(ctx, err, codeByErrPOST)
		return handler_errors.InvalidParameters
	}
	h.Respond(ctx, http.StatusCreated, plan_dto.ToResponseAcceptPlan(res, req.Skills))
	return nil
}

func (h *AcceptHandler) DELETE(ctx echo.Context) error {
	userID, ok := ctx.Request().Context().Value("user_id").(int64)
	if !ok {
		h.Log(ctx.Request()).Error("can not get user_id from context")
		h.Error(ctx, http.StatusInternalServerError, handler_errors.InternalError)
		return handler_errors.InternalError
	}
	offerID, _, err := h.GetInt64FromParam(ctx, "id")
	if err != nil {
		h.Log(ctx.Request()).Warnf("OfferIDAcceptHandler: can not parse path param ID %s", err)
		h.Error(ctx, http.StatusBadRequest, handler_errors.InvalidParameters)
		return handler_errors.InvalidParameters
	}

	if err := h.offerUsecase.Delete(userID, offerID); err != nil {
		h.Log(ctx.Request()).Warnf("OfferIDAcceptHandler: can not delete offer: userID = %v, offerID = %v", userID, offerID)
		h.UsecaseError(ctx, err, codeByErrDELETE)
		return err
	}
	ctx.Response().WriteHeader(http.StatusOK)
	return nil
}
