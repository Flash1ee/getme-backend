package user_telegram_auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"getme-backend/internal/app/middleware"
	token_usecase "getme-backend/internal/app/token/usecase"
	"getme-backend/internal/app/user/dto"
	"getme-backend/internal/app/user/usecase"
	"getme-backend/internal/microservices/auth/delivery/grpc/client"
	middleware2 "getme-backend/internal/microservices/auth/sessions/middleware"
	"getme-backend/internal/pkg/adapter/echo_adapter"
	bh "getme-backend/internal/pkg/handler"
	"getme-backend/internal/pkg/handler/handler_errors"
)

type UserAuthCheckHandler struct {
	userUsecase   user_usecase.Usecase
	tokenUsecase  token_usecase.Usecase
	sessionClient client.AuthCheckerClient
	bh.BaseHandler
}

func NewUserAuthCheckHandler(log *logrus.Logger, ucUser user_usecase.Usecase, sessionClient client.AuthCheckerClient) *UserAuthCheckHandler {
	h := &UserAuthCheckHandler{
		BaseHandler:   *bh.NewBaseHandler(log),
		userUsecase:   ucUser,
		sessionClient: sessionClient,
	}
	h.AddMiddleware(echo_adapter.WrapMiddleware(middleware.NewUtilitiesMiddleware(logrus.New()).CheckPanic))
	h.AddMethod(http.MethodGet, h.GET, echo_adapter.WrapMiddlewareToFunc(middleware2.NewSessionMiddleware(h.sessionClient, log).CheckNotAuthorized))
	return h
}

func (h *UserAuthCheckHandler) GET(ctx echo.Context) error {
	req := &dto.UserAuthTelegramCheckRequest{}

	_, status := h.GetParamToStruct(ctx, req)
	if status != bh.OK {
		h.Log(ctx.Request()).Errorf("AUTH_CHECK Handler: Error get params fro auth request %v\n", req)
		ctx.Response().WriteHeader(http.StatusBadRequest)
		return nil
	}

	u, err := h.userUsecase.AuthTelegram(req.ToUserAuthUsecase())
	if err != nil {
		h.Log(ctx.Request()).Warnf("error auth usecase; %v, req data: %v", err, req)
		h.UsecaseError(ctx, err, codesByErrorsGET)
		return err
	}

	res, err := h.sessionClient.CreateByToken(ctx.Request().Context(), req.Token, u.ID)
	if err != nil || res.TokenID != req.Token {
		h.Log(ctx.Request()).Errorf("Error create session %s", err)
		h.Error(ctx, http.StatusInternalServerError, handler_errors.ErrorCreateSession)
		return err
	}

	h.Log(ctx.Request()).Debugf("user success auth %v", u)
	h.Respond(ctx, http.StatusCreated, dto.ToUserResponseFromUsecase(u))
	return nil
}
