package simple_auth_handler

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"getme-backend/internal/app/auth/dto"
	auth_usecase "getme-backend/internal/app/auth/usecase"
	session_client "getme-backend/internal/microservices/auth/delivery/grpc/client"
	"getme-backend/internal/microservices/auth/sessions/middleware"
	"getme-backend/internal/microservices/auth/sessions/usecase"
	"getme-backend/internal/pkg/adapter/echo_adapter"
	bh "getme-backend/internal/pkg/handler"
	"getme-backend/internal/pkg/handler/handler_errors"

	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	sessionClient session_client.AuthCheckerClient
	authUsecase   auth_usecase.Usecase
	bh.BaseHandler
}

func NewAuthHandler(log *logrus.Logger, sClient session_client.AuthCheckerClient,
	ucAuth auth_usecase.Usecase) *AuthHandler {
	h := &AuthHandler{
		BaseHandler:   *bh.NewBaseHandler(log),
		sessionClient: sClient,
		authUsecase:   ucAuth,
	}
	h.AddMiddleware(echo_adapter.WrapMiddleware(middleware.NewSessionMiddleware(h.sessionClient, log).CheckNotAuthorized))
	h.AddMethod(http.MethodPost, h.POST)
	return h
}

// POST Login
// @Summary login user
// @Description login user
// @tags user
// @Accept  json
// @Produce json
// @Param user body http_models.RequestLogin true "Request body for login"
// @Success 200 "Successfully login"
// @Failure 404 {object} http_models.ErrResponse "user not found"
// @Failure 422 {object} http_models.ErrResponse "invalid body in request"
// @Failure 500 {object} http_models.ErrResponse "can not create session", "can not do bd operation"
// @Failure 401 {object} http_models.ErrResponse "incorrect email or password"
// @Failure 418 "User are authorized"
// @Router /login [POST]
func (h *AuthHandler) POST(ctx echo.Context) error {
	req := &dto.AuthSimpleRequest{}
	err := h.GetRequestBody(ctx, req)
	if err != nil {
		h.Log(ctx.Request()).Warnf("can not decode body %s", err)
		h.Error(ctx, http.StatusUnprocessableEntity, handler_errors.InvalidBody)
		return handler_errors.InvalidBody
	}
	h.Log(ctx.Request()).Debugf("Login : %s, password : %s", req.Login, req.Password)

	id, err := h.authUsecase.AuthSimple(req.Login, req.Password)
	if err != nil {
		h.UsecaseError(ctx, err, codesByErrors)
		return err
	}

	res, err := h.sessionClient.Create(context.Background(), id)
	if err != nil || res.UserID != id {
		h.Log(ctx.Request()).Errorf("Error create session %s", err)
		h.Error(ctx, http.StatusInternalServerError, handler_errors.ErrorCreateSession)
		return handler_errors.ErrorCreateSession
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    res.UniqID,
		Path:     "/",
		Expires:  time.Now().Add(usecase.ExpiredCookiesTime),
		HttpOnly: true,
	}

	http.SetCookie(ctx.Response(), cookie)
	ctx.Response().WriteHeader(http.StatusOK)
	return nil
}
