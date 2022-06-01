package telegram_register_handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"getme-backend/internal/app/auth/dto"
	"getme-backend/internal/app/auth/usecase"
	"getme-backend/internal/app/middleware"
	token_usecase "getme-backend/internal/app/token/usecase"
	user_usecase "getme-backend/internal/app/user/usecase"
	"getme-backend/internal/microservices/auth/delivery/grpc/client"
	middleware2 "getme-backend/internal/microservices/auth/sessions/middleware"
	"getme-backend/internal/pkg/adapter/echo_adapter"
	bh "getme-backend/internal/pkg/handler"
	"getme-backend/internal/pkg/handler/handler_errors"
)

type RegisterHandler struct {
	authUsecase auth_usecase.Usecase
	userUsecase user_usecase.Usecase

	tokenUsecase  token_usecase.Usecase
	sessionClient client.AuthCheckerClient
	bh.BaseHandler
}

func NewRegisterHandler(log *logrus.Logger, sessionClient client.AuthCheckerClient,
	ucAuth auth_usecase.Usecase, ucUser user_usecase.Usecase) *RegisterHandler {
	h := &RegisterHandler{
		BaseHandler:   *bh.NewBaseHandler(log),
		authUsecase:   ucAuth,
		userUsecase:   ucUser,
		sessionClient: sessionClient,
	}
	h.AddMiddleware(echo_adapter.WrapMiddleware(middleware.NewUtilitiesMiddleware(logrus.New()).CheckPanic))
	h.AddMethod(http.MethodGet, h.GET, echo_adapter.WrapMiddlewareToFunc(middleware2.NewSessionMiddleware(h.sessionClient, log).CheckNotAuthorized))
	return h
}

func (h *RegisterHandler) GET(ctx echo.Context) error {
	req := &dto.AuthTelegramCheckRequest{}

	_, status := h.GetParamToStruct(ctx, req)
	if status != bh.OK {
		h.Log(ctx.Request()).Errorf("AUTH_CHECK Handler: Error get params fro auth request %v\n", req)
		ctx.Response().WriteHeader(http.StatusBadRequest)
		return nil
	}
	userID := int64(-1)
	id, err := h.userUsecase.FindByNickname(req.Username)
	if err != nil {
		if err != user_usecase.UserNotFound {
			h.Log(ctx.Request()).Warnf("RegisterHandler - Auth: FindByNickname(%v) error -  %s", req.Username, err)
			h.Error(ctx, http.StatusInternalServerError, handler_errors.InternalError)
			return handler_errors.InternalError
		}
		userID, err = h.userUsecase.CreateFilledUser(req.ToUserUsecase())
		if err != nil {
			h.Log(ctx.Request()).Warnf("RegisterHandler - AuthTelegram: can not create Filled user, err %s", err)
			h.Error(ctx, http.StatusInternalServerError, handler_errors.InternalError)
			return handler_errors.InternalError
		}
	} else {
		userID = id.ID.Int64
	}

	usecaseDTO := req.ToAuthUsecase()
	usecaseDTO.ID = userID
	u, err := h.authUsecase.AuthTelegram(usecaseDTO)
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

	//queryParam := url.Values{
	//	"token": []string{res.TokenID},
	//}
	//ctx.Redirect(http.StatusPermanentRedirect, "/api/v1/auth/telegram/login"+"?"+queryParam.Encode())
	h.Respond(ctx, http.StatusCreated, dto.ToUserResponseFromUsecase(u))
	return nil
}
