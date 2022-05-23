package user_telegram_auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"getme-backend/internal/app/middleware"
	token_usecase "getme-backend/internal/app/token/usecase"
	"getme-backend/internal/app/user/usecase"
	"getme-backend/internal/microservices/auth/delivery/grpc/client"
	middleware2 "getme-backend/internal/microservices/auth/sessions/middleware"
	"getme-backend/internal/pkg/adapter/echo_adapter"
	bh "getme-backend/internal/pkg/handler"
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
	return nil

}
