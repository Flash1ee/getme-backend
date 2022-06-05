package token_handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"getme-backend/internal/app/middleware"
	"getme-backend/internal/app/token/dto"
	token_usecase "getme-backend/internal/app/token/usecase"
	"getme-backend/internal/microservices/auth/delivery/grpc/client"
	middleware2 "getme-backend/internal/microservices/auth/sessions/middleware"
	"getme-backend/internal/pkg/adapter/echo_adapter"
	bh "getme-backend/internal/pkg/handler"
)

type TokenHandler struct {
	tokenUsecase  token_usecase.Usecase
	sessionClient client.AuthCheckerClient
	bh.BaseHandler
}

func NewTokenHandler(log *logrus.Logger, tokenUc token_usecase.Usecase, sessionClient client.AuthCheckerClient) *TokenHandler {
	h := &TokenHandler{
		BaseHandler:   *bh.NewBaseHandler(log),
		tokenUsecase:  tokenUc,
		sessionClient: sessionClient,
	}
	h.AddMiddleware(echo_adapter.WrapMiddleware(middleware.NewUtilitiesMiddleware(logrus.New()).CheckPanic))
	h.AddMethod(http.MethodGet, h.GET, echo_adapter.WrapMiddlewareToFunc(middleware2.NewSessionMiddleware(h.sessionClient, log).CheckNotAuthorized))
	return h
}

func (h *TokenHandler) GET(ctx echo.Context) error {
	tokenSources := dto.TokenSourcesUsecase{
		IdentifierData: ctx.Request().RemoteAddr,
	}
	token, err := h.tokenUsecase.GetTokenByData(tokenSources)

	if err != nil {
		h.Log(ctx.Request()).Warnf("TOKEN HANDLER: error token usecase; %v\n", err)
		h.UsecaseError(ctx, err, codeByErrorGET)
		return err
	}

	h.Log(ctx.Request()).Debugf("get token success, token = %v", token.Token)
	h.Respond(ctx, http.StatusOK, token)
	return nil
}
