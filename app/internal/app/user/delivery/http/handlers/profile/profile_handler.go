package user_profile_handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"getme-backend/internal/app/user/dto"
	user_usecase "getme-backend/internal/app/user/usecase"
	session_client "getme-backend/internal/microservices/auth/delivery/grpc/client"
	"getme-backend/internal/microservices/auth/sessions/middleware"
	"getme-backend/internal/pkg/adapter/echo_adapter"
	bh "getme-backend/internal/pkg/handler"
	"getme-backend/internal/pkg/handler/handler_errors"
)

type ProfileHandler struct {
	sessionClient session_client.AuthCheckerClient
	userUsecase   user_usecase.Usecase
	bh.BaseHandler
}

func NewProfileHandler(log *logrus.Logger, sClient session_client.AuthCheckerClient,
	ucUser user_usecase.Usecase) *ProfileHandler {
	h := &ProfileHandler{
		BaseHandler:   *bh.NewBaseHandler(log),
		sessionClient: sClient,
		userUsecase:   ucUser,
	}
	h.AddMiddleware(echo_adapter.WrapMiddleware(middleware.NewSessionMiddleware(h.sessionClient, log).Check))
	h.AddMiddleware(echo_adapter.WrapMiddleware(middleware.NewSessionMiddleware(h.sessionClient, log).AddUserId))

	h.AddMethod(http.MethodGet, h.GET)
	return h
}

func (h *ProfileHandler) GET(ctx echo.Context) error {
	userID, ok := ctx.Request().Context().Value("user_id").(int64)
	if !ok {
		h.Log(ctx.Request()).Error("can not get user_id from context")
		h.Error(ctx, http.StatusInternalServerError, handler_errors.InternalError)
		return handler_errors.InternalError
	}
	res, err := h.userUsecase.FindByID(userID)
	if err != nil {
		h.UsecaseError(ctx, err, codesByErrorsGET)
		return err
	}
	h.Respond(ctx, http.StatusOK, dto.ToUserResponse(res))
	return nil
}
