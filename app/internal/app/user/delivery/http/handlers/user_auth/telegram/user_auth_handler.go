package telegram

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"getme-backend/internal/app/middleware"
	dto_token "getme-backend/internal/app/token/dto"
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

func NewUserAuthHandler(log *logrus.Logger, ucUser user_usecase.Usecase, sessionClient client.AuthCheckerClient, ucToken token_usecase.Usecase) *UserAuthHandler {
	h := &UserAuthHandler{
		BaseHandler:   *bh.NewBaseHandler(log),
		userUsecase:   ucUser,
		tokenUsecase:  ucToken,
		sessionClient: sessionClient,
	}
	h.AddMiddleware(echo_adapter.WrapMiddleware(middleware.NewUtilitiesMiddleware(logrus.New()).CheckPanic))
	h.AddMethod(http.MethodGet, h.GET, echo_adapter.WrapMiddlewareToFunc(middleware2.NewSessionMiddleware(h.sessionClient, log).CheckNotAuthorized))
	return h
}
func (h *UserAuthHandler) getQueryParams(ctx echo.Context) (*dto.UserAuthRequest, int, error) {
	req := dto.NewUserAuthRequest()
	binder := echo.QueryParamsBinder(ctx)

	errs := binder.String("token", &req.Token).
		BindErrors()

	if errs != nil {
		for _, err := range errs {
			bErr := err.(*echo.BindingError)
			h.Log(ctx.Request()).Errorf("AUTH HANDLER: error get query param with tag field %v value = %v\n", bErr.Field, bErr.Values)
		}
		ctx.Response().WriteHeader(http.StatusBadRequest)
		return nil, len(errs), errs[0]
	}

	if err := h.Validator.Struct(req); err != nil {
		h.Log(ctx.Request()).Errorf("AUTH HANDLER: validate error, req = %v err = %v\n", req, err)
		return nil, 1, err

	}
	return req, 0, nil
}

func (h *UserAuthHandler) GET(ctx echo.Context) error {
	req, errsCount, err := h.getQueryParams(ctx)
	if errsCount != 0 {
		h.Log(ctx.Request()).Errorf("AUTH HANDLER: Error validation request params %v\n", req)
		h.Error(ctx, http.StatusBadRequest, handler_errors.InvalidQueries)
		return handler_errors.InvalidQueries
	}
	tokenSources := dto_token.TokenSourcesUsecase{
		IdentifierData: ctx.Request().RemoteAddr,
	}
	tokenDTO := dto_token.TokenUsecase{
		Token: req.Token,
	}

	if err := h.tokenUsecase.Check(tokenSources, tokenDTO); err != nil {
		h.Log(ctx.Request()).Warnf(
			"AUTH HANDLER: invalid token, not correct validation; %v, req data - token: %v, identifierData = %v",
			err, req, tokenSources.IdentifierData)

		h.UsecaseError(ctx, err, codesByErrorsGET)
		return err
	}
	u, err := h.sessionClient.CheckWithDelete(ctx.Request().Context(), req.Token)
	if err != nil {
		h.Log(ctx.Request()).Warnf("AUTH HANDLER: error check token; %v, req data - token: %v", err, req)
		h.UsecaseError(ctx, err, codesByErrorsGET)
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
