package user_by_id_handler

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

type UserByIDHandler struct {
	sessionClient session_client.AuthCheckerClient
	userUsecase   user_usecase.Usecase
	bh.BaseHandler
}

func NewUserByIDHandler(log *logrus.Logger, sClient session_client.AuthCheckerClient,
	ucUser user_usecase.Usecase) *UserByIDHandler {
	h := &UserByIDHandler{
		BaseHandler:   *bh.NewBaseHandler(log),
		sessionClient: sClient,
		userUsecase:   ucUser,
	}
	h.AddMiddleware(echo_adapter.WrapMiddleware(middleware.NewSessionMiddleware(h.sessionClient, log).Check))
	h.AddMiddleware(echo_adapter.WrapMiddleware(middleware.NewSessionMiddleware(h.sessionClient, log).AddUserId))

	h.AddMethod(http.MethodGet, h.GET)

	return h
}
func (h *UserByIDHandler) GET(ctx echo.Context) error {
	userID, status, err := h.GetInt64FromParam(ctx, "id")
	if err != nil {
		h.Error(ctx, status, err)
		return err
	}
	if len(ctx.ParamValues()) > 1 {
		h.Log(ctx.Request()).Warnf("Too many parametres %v", ctx.ParamValues())
		h.Error(ctx, http.StatusBadRequest, handler_errors.InvalidParameters)
		return handler_errors.InvalidParameters
	}
	res, err := h.userUsecase.FindByID(userID)
	if err != nil {
		h.UsecaseError(ctx, err, codeByErrorGET)
		return err
	}
	h.Respond(ctx, http.StatusOK, dto.ToUserResponse(res))
	return nil
}
