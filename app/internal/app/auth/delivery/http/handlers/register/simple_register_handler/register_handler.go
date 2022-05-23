package simple_register_handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"getme-backend/internal/app/auth/dto"
	auth_usecase "getme-backend/internal/app/auth/usecase"
	user_usecase "getme-backend/internal/app/user/usecase"
	session_client "getme-backend/internal/microservices/auth/delivery/grpc/client"
	bh "getme-backend/internal/pkg/handler"
	"getme-backend/internal/pkg/handler/handler_errors"
)

type RegisterHandler struct {
	sessionClient session_client.AuthCheckerClient
	authUsecase   auth_usecase.Usecase
	userUsecase   user_usecase.Usecase
	bh.BaseHandler
}

func NewRegisterHandler(log *logrus.Logger, sClient session_client.AuthCheckerClient,
	ucUser user_usecase.Usecase, ucAuth auth_usecase.Usecase) *RegisterHandler {
	h := &RegisterHandler{
		sessionClient: sClient,
		authUsecase:   ucAuth,
		userUsecase:   ucUser,
		BaseHandler:   *bh.NewBaseHandler(log),
	}
	h.AddMethod(http.MethodPost, h.POST)
	return h
}

// POST Registration
// @Summary create new user
// @tags user
// @Description create new account and get cookies
// @Accept  json
// @Produce json
// @Param register_info body http_models.RequestRegistration true "Request body for user registration"
// @Success 201 {object} http_models.IdResponse "CreateSimple user successfully"
// @Failure 409 {object} http_models.ErrResponse "user already exist"
// @Failure 422 {object} http_models.ErrResponse "invalid body in request", "nickname already exist", "incorrect email or password", "incorrect nickname"
// @Failure 500 {object} http_models.ErrResponse "can not do bd operation"
// @Failure 418 "User are authorized"
// @Router /register [POST]
func (h *RegisterHandler) POST(ctx echo.Context) error {
	req := &dto.AuthSimpleRegistrationRequest{}

	err := h.GetRequestBody(ctx, req)
	if err != nil {
		h.Log(ctx.Request()).Warnf("RegisterHandler - Auth: can not parse request %s", err)
		h.Error(ctx, http.StatusUnprocessableEntity, handler_errors.InvalidBody)
		return handler_errors.InvalidBody
	}
	var userID int64
	_, err = h.userUsecase.FindByNickname(req.Login)
	if err != nil {
		if errors.Is(err, user_usecase.UserNotFound) {
			userID, err = h.userUsecase.CreateBaseUser(req.Login)
			if err != nil {
				h.Log(ctx.Request()).Warnf("RegisterHandler - Auth: can not create Base user, err %s", err)
				h.Error(ctx, http.StatusInternalServerError, handler_errors.InternalError)
				return handler_errors.InternalError
			}
		} else {
			h.Log(ctx.Request()).Warnf("RegisterHandler - Auth: FindByNickname(%v, %v) error -  %s", req.Login, req.Password, err)
			h.Error(ctx, http.StatusInternalServerError, handler_errors.InternalError)
			return handler_errors.InternalError
		}

	}

	usecaseDTO := req.ToSimpleRegistrationUsecase()
	usecaseDTO.UserID = userID

	id, err := h.authUsecase.CreateSimple(usecaseDTO)
	if err != nil {
		h.UsecaseError(ctx, err, codeByError)
		return err
	}

	h.Respond(ctx, http.StatusCreated, dto.IDResponse{ID: id})
	return nil
}
