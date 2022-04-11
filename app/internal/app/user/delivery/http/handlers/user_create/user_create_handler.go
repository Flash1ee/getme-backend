package user_create_handler

import (
	"net/http"

	"github.com/qiangxue/fasthttp-routing"
	"github.com/sirupsen/logrus"

	"getme-backend/internal/app/user/dto"
	"getme-backend/internal/app/user/entities"
	"getme-backend/internal/app/user/repository"
	bh "getme-backend/internal/pkg/handler"
	"getme-backend/internal/pkg/handler/handler_errors"
	"getme-backend/internal/pkg/utilits/postgresql"
)

type UserCreateHandler struct {
	userRepository repository.Repository
	bh.BaseHandler
}

func NewUserCreateHandler(log *logrus.Logger, rep repository.Repository) *UserCreateHandler {
	h := &UserCreateHandler{
		BaseHandler:    *bh.NewBaseHandler(log),
		userRepository: rep,
	}
	h.AddMethod(http.MethodPost, h.POST)
	return h
}

func (h *UserCreateHandler) POST(ctx *routing.Context) error {
	req := &dto.UserRequest{}
	err := h.GetRequestBody(ctx, req)
	if err != nil {
		h.Log(ctx).Warnf("can not parse request %s", err)
		h.Error(ctx, http.StatusUnprocessableEntity, handler_errors.InvalidBody)
		return nil
	}

	nickname, status := h.GetStringFromParam(ctx, "nickname")
	if status == bh.EmptyQuery {
		ctx.SetStatusCode(http.StatusBadRequest)
		return nil
	}

	u, err := h.userRepository.Create(&entities.User{
		Nickname: nickname,
		About:    req.About,
		Email:    req.Email,
		Fullname: req.Fullname,
	})

	if err == postgresql_utilits.Conflict {
		h.Log(ctx).Warnf("conflict with request %v", req)
		h.Respond(ctx, http.StatusConflict, dto.ToUsersResponse(u))
		return nil
	}

	if err != nil {
		h.UsecaseError(ctx, err, codesByErrorsPOST)
		return nil
	}

	h.Log(ctx).Debugf("create user %v", u)
	h.Respond(ctx, http.StatusCreated, dto.ToUserResponse(&u[0]))
	return nil
}
