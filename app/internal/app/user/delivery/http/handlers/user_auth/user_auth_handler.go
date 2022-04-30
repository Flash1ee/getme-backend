package user_auth_handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"getme-backend/internal/app/user/usecase"

	"getme-backend/internal/app/middleware"
	"getme-backend/internal/app/user/dto"
	"getme-backend/internal/pkg/adapter/echo_adapter"
	bh "getme-backend/internal/pkg/handler"
)

type UserAuthHandler struct {
	userUsecase user_usecase.Usecase
	bh.BaseHandler
}

func NewUserAuthHandler(log *logrus.Logger, uc user_usecase.Usecase) *UserAuthHandler {
	h := &UserAuthHandler{
		BaseHandler: *bh.NewBaseHandler(log),
		userUsecase: uc,
	}
	h.AddMiddleware(echo_adapter.WrapMiddleware(middleware.NewUtilitiesMiddleware(logrus.New()).CheckPanic))
	h.AddMethod(http.MethodGet, h.GET)
	return h
}

func (h *UserAuthHandler) GET(ctx echo.Context) error {
	req := &dto.UserAuthRequest{}

	_, status := h.GetParamToStruct(ctx, req)
	if status == bh.EmptyQuery {
		ctx.Response().WriteHeader(http.StatusBadRequest)
		return nil
	}

	u, err := h.userUsecase.Auth(req.ToUserAuthUsecase())

	if err != nil {
		h.Log(ctx.Request()).Warnf("error auth usecase; %v, req data: %v", err, req)
		h.UsecaseError(ctx, err, codesByErrorsPOST)
		return nil
	}
	ctx.HTML(200, "<!DOCTYPE html>\n<html>\n    "+
		"<head>\n        <title>Example</title>\n    </head>\n    "+
		"<body>\n        <p>This is an example of a simple HTML page with one paragraph.</p>\n    "+
		"</body>\n</html>")

	h.Log(ctx.Request()).Debugf("user success auth %v", u)
	h.Respond(ctx, http.StatusCreated, u)
	return nil
}
