package user_logout

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	session_client "getme-backend/internal/microservices/auth/delivery/grpc/client"
	session_middleware "getme-backend/internal/microservices/auth/sessions/middleware"
	"getme-backend/internal/pkg/adapter/echo_adapter"
	bh "getme-backend/internal/pkg/handler"
	"getme-backend/internal/pkg/handler/handler_errors"

	"github.com/sirupsen/logrus"

	"getme-backend/internal/microservices/auth/delivery/grpc/client"
)

type LogoutHandler struct {
	sessionClient client.AuthCheckerClient
	bh.BaseHandler
}

func NewLogoutHandler(log *logrus.Logger,
	sClient session_client.AuthCheckerClient) *LogoutHandler {
	h := &LogoutHandler{
		BaseHandler:   *bh.NewBaseHandler(log),
		sessionClient: sClient,
	}
	h.AddMethod(http.MethodPost, h.POST,
		echo_adapter.WrapMiddlewareToFunc(session_middleware.NewSessionMiddleware(h.sessionClient, log).CheckFunc),
	)

	return h
}

// POST Logout
// @Summary logout user
// @Description logout user
// @tags user
// @Accept  json
// @Produce json
// @Success 201 "Successfully logout"
// @Failure 500 {object} http_models.ErrResponse "server error"
// @Failure 401 "User not are authorized"
// @Router /logout [POST]
func (h *LogoutHandler) POST(ctx echo.Context) error {
	uniqID := ctx.Request().Context().Value("session_id")
	if uniqID == nil {
		h.Log(ctx.Request()).Error("can not get session_id from context")
		h.Error(ctx, http.StatusInternalServerError, handler_errors.InternalError)
		return handler_errors.InternalError
	}

	h.Log(ctx.Request()).Debugf("Logout session: %s", uniqID)

	err := h.sessionClient.Delete(context.Background(), uniqID.(string))
	if err != nil {
		h.Log(ctx.Request()).Errorf("can not delete session %s", err)
		h.Error(ctx, http.StatusInternalServerError, handler_errors.DeleteCookieFail)
		return handler_errors.DeleteCookieFail
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    uniqID.(string),
		Path:     "/",
		Expires:  time.Now().AddDate(0, 0, -1),
		HttpOnly: true,
	}
	http.SetCookie(ctx.Response(), cookie)
	ctx.Response().WriteHeader(http.StatusOK)
	return nil
}
