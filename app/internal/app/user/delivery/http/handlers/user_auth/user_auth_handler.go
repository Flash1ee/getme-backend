package user_auth_handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"getme-backend/internal/app/middleware"
	dto2 "getme-backend/internal/app/token/dto"
	token_usecase "getme-backend/internal/app/token/usecase"
	"getme-backend/internal/app/user/dto"
	user_usecase "getme-backend/internal/app/user/usecase"
	"getme-backend/internal/microservices/auth/delivery/grpc/client"
	middleware2 "getme-backend/internal/microservices/auth/sessions/middleware"
	"getme-backend/internal/microservices/auth/sessions/usecase"
	"getme-backend/internal/pkg/adapter/echo_adapter"
	bh "getme-backend/internal/pkg/handler"
	"getme-backend/internal/pkg/handler/handler_errors"
)

type UserAuthHandler struct {
	userUsecase   user_usecase.Usecase
	tokenUsecase  token_usecase.Usecase
	sessionClient client.AuthCheckerClient
	bh.BaseHandler
}

func NewUserAuthHandler(log *logrus.Logger, ucUser user_usecase.Usecase, sessionClient client.AuthCheckerClient) *UserAuthHandler {
	h := &UserAuthHandler{
		BaseHandler:   *bh.NewBaseHandler(log),
		userUsecase:   ucUser,
		sessionClient: sessionClient,
	}
	h.AddMiddleware(echo_adapter.WrapMiddleware(middleware.NewUtilitiesMiddleware(logrus.New()).CheckPanic))
	h.AddMethod(http.MethodGet, h.GET, echo_adapter.WrapMiddlewareToFunc(middleware2.NewSessionMiddleware(h.sessionClient, log).CheckNotAuthorized))
	return h
}

func (h *UserAuthHandler) GET(ctx echo.Context) error {
	req := &dto.UserAuthRequest{}

	_, status := h.GetParamToStruct(ctx, req)
	if status != bh.OK {
		h.Log(ctx.Request()).Errorf("AUTH HANDLER: Error get params fro auth request %v\n", req)
		ctx.Response().WriteHeader(http.StatusBadRequest)
		return nil
	}

	tokenSources := dto2.TokenSourcesUsecase{
		IdentifierData: ctx.Request().RemoteAddr,
	}
	tokenDTO := dto2.TokenUsecase{
		Token: req.Token,
	}

	if err := h.tokenUsecase.Check(tokenSources, tokenDTO); err != nil {
		h.Log(ctx.Request()).Warnf(
			"AUTH HANDLER: invalid token, not correct validation; %v, req data - token: %v, identifierData = %v",
			err, req, tokenSources.IdentifierData)

		h.UsecaseError(ctx, err, codesByErrorsPOST)
		return err
	}
	u, err := h.sessionClient.CheckWithDelete(ctx.Request().Context(), req.Token)
	if err != nil {
		h.Log(ctx.Request()).Warnf("AUTH HANDLER: error check token; %v, req data - token: %v", err, req)
		h.UsecaseError(ctx, err, codesByErrorsPOST)
		return err
	}
	userID, err := strconv.Atoi(u.UserID)
	if err != nil {
		h.Log(ctx.Request()).Warnf("AUTH HANDLER: error convert userID to int. userID; %v, err: %v", userID, err)
		h.UsecaseError(ctx, err, codesByErrorsGET)
		return err
	}

	res, err := h.sessionClient.Create(ctx.Request().Context(), int64(userID))
	if err != nil || res.UserID != int64(userID) {
		h.Log(ctx.Request()).Errorf("AUTH HANDLER: error create session %s", err)
		h.Error(ctx, http.StatusInternalServerError, handler_errors.ErrorCreateSession)
		return err
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    res.UniqID,
		Path:     "/",
		Expires:  time.Now().Add(usecase.ExpiredCookiesTime),
		HttpOnly: true,
	}

	ctx.SetCookie(cookie)
	h.Log(ctx.Request()).Debugf("user success auth %v", u)
	ctx.Response().WriteHeader(http.StatusOK)
	return nil
}
