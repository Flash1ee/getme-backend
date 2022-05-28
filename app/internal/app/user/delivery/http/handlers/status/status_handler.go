package user_status_handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"getme-backend/internal/app/user/dto"
	"getme-backend/internal/app/user/usecase"
	session_client "getme-backend/internal/microservices/auth/delivery/grpc/client"
	"getme-backend/internal/microservices/auth/sessions/middleware"
	"getme-backend/internal/pkg/adapter/echo_adapter"
	bh "getme-backend/internal/pkg/handler"
	"getme-backend/internal/pkg/handler/handler_errors"
)

type StatusHandler struct {
	sessionClient session_client.AuthCheckerClient
	userUsecase   user_usecase.Usecase
	bh.BaseHandler
}

func NewStatusHandler(log *logrus.Logger, sClient session_client.AuthCheckerClient,
	ucUser user_usecase.Usecase) *StatusHandler {
	h := &StatusHandler{
		BaseHandler:   *bh.NewBaseHandler(log),
		sessionClient: sClient,
		userUsecase:   ucUser,
	}
	h.AddMiddleware(echo_adapter.WrapMiddleware(middleware.NewSessionMiddleware(h.sessionClient, log).Check))
	h.AddMiddleware(echo_adapter.WrapMiddleware(middleware.NewSessionMiddleware(h.sessionClient, log).AddUserId))

	h.AddMethod(http.MethodGet, h.GET)
	h.AddMethod(http.MethodPut, h.PUT)

	return h
}
func (h *StatusHandler) GET(ctx echo.Context) error {
	userID, ok := ctx.Request().Context().Value("user_id").(int64)
	if !ok {
		h.Log(ctx.Request()).Error("can not get user_id from context")
		h.Error(ctx, http.StatusInternalServerError, handler_errors.InternalError)
		return handler_errors.InternalError
	}
	res, err := h.userUsecase.GetMentorStatus(userID)
	if err != nil {
		h.UsecaseError(ctx, err, codesByErrorsGET)
		return err
	}

	h.Respond(ctx, http.StatusOK, res.ToResponseStatus())
	return nil
}
func (h *StatusHandler) PUT(ctx echo.Context) error {
	userID, ok := ctx.Request().Context().Value("user_id").(int64)
	if !ok {
		h.Log(ctx.Request()).Error("can not get user_id from context")
		h.Error(ctx, http.StatusInternalServerError, handler_errors.InternalError)
		return handler_errors.InternalError
	}

	res, err := h.userUsecase.UpdateMentorStatus(&dto.UserStatusUsecase{
		UserID: userID,
	})
	if err != nil {
		h.UsecaseError(ctx, err, codesByErrorsPUT)
		return err
	}

	h.Respond(ctx, http.StatusOK, res.ToResponseStatus())
	return nil
}
