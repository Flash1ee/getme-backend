package offer_handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"getme-backend/internal/app/offer/dto"
	offer_usecase "getme-backend/internal/app/offer/usecase"
	dto2 "getme-backend/internal/app/user/dto"
	session_client "getme-backend/internal/microservices/auth/delivery/grpc/client"
	middleware2 "getme-backend/internal/microservices/auth/sessions/middleware"
	"getme-backend/internal/pkg/adapter/echo_adapter"
	bh "getme-backend/internal/pkg/handler"
	"getme-backend/internal/pkg/handler/handler_errors"
)

type OfferHandler struct {
	sessionClient session_client.AuthCheckerClient
	offerUsecase  offer_usecase.Usecase
	bh.BaseHandler
}

func NewOfferHandler(log *logrus.Logger, sClient session_client.AuthCheckerClient, offerUs offer_usecase.Usecase) *OfferHandler {
	h := &OfferHandler{
		BaseHandler:   *bh.NewBaseHandler(log),
		sessionClient: sClient,
		offerUsecase:  offerUs,
	}

	h.AddMethod(http.MethodGet, h.POST, echo_adapter.WrapMiddlewareToFunc(middleware2.NewSessionMiddleware(sClient, log).CheckFunc))
	h.AddMethod(http.MethodPost, h.GET, echo_adapter.WrapMiddlewareToFunc(middleware2.NewSessionMiddleware(sClient, log).CheckFunc))

	return h
}

func (h *OfferHandler) GET(ctx echo.Context) error {
	userID, ok := ctx.Request().Context().Value("user_id").(int64)
	if !ok {
		h.Log(ctx.Request()).Error("can not get user_id from context")
		h.Error(ctx, http.StatusInternalServerError, handler_errors.InternalError)
		return handler_errors.InternalError
	}
	res, err := h.offerUsecase.Get(userID)
	if err != nil {
		h.UsecaseError(ctx, err, codeByErrGET)
		return err
	}

	h.Respond(ctx, http.StatusOK, dto2.ToUsersResponse(res))
	return nil
}

func (h *OfferHandler) POST(ctx echo.Context) error {
	userID, ok := ctx.Request().Context().Value("user_id").(int64)
	if !ok {
		h.Log(ctx.Request()).Error("can not get user_id from context")
		h.Error(ctx, http.StatusInternalServerError, handler_errors.InternalError)
		return handler_errors.InternalError
	}
	req := &dto.RequestCreateOffer{}

	err := h.GetRequestBody(ctx, req)
	if err != nil {
		h.Log(ctx.Request()).Warnf("can not parse request %s", err)
		h.Error(ctx, http.StatusUnprocessableEntity, handler_errors.InvalidBody)
		return handler_errors.InvalidBody
	}
	usecaseDTO := dto.ToOfferUsecaseDTO(*req)
	usecaseDTO.MenteeID = userID

	res, err := h.offerUsecase.Create(usecaseDTO)
	if err != nil {
		h.UsecaseError(ctx, err, codeByErrPOST)
		return err
	}
	h.Respond(ctx, http.StatusOK, dto.ResponseOffer{
		ID: res,
	})
	return nil
}
