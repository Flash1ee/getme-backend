package user_profile_handler

import (
	"net/http"

	"github.com/qiangxue/fasthttp-routing"
	"github.com/sirupsen/logrus"

	"getme-backend/internal/app/user/dto"
	"getme-backend/internal/app/user/entities"
	"getme-backend/internal/app/user/repository"
	bh "getme-backend/internal/pkg/handler"
	"getme-backend/internal/pkg/handler/handler_errors"
)

type UserProfileHandler struct {
	userRepository repository.Repository
	bh.BaseHandler
}

func NewUserProfileHandler(log *logrus.Logger, rep repository.Repository) *UserProfileHandler {
	h := &UserProfileHandler{
		BaseHandler:    *bh.NewBaseHandler(log),
		userRepository: rep,
	}
	h.AddMethod(http.MethodGet, h.GET)
	h.AddMethod(http.MethodPost, h.POST)
	return h
}

func (h *UserProfileHandler) GET(ctx *routing.Context) error {
	nickname, status := h.GetStringFromParam(ctx, "nickname")
	if status == bh.EmptyQuery {
		ctx.SetStatusCode(http.StatusBadRequest)
		return nil
	}

	u, err := h.userRepository.Get(nickname)
	if err != nil {
		h.UsecaseError(ctx, err, codesByErrorsGET)
		return nil
	}

	h.Log(ctx).Debugf("get user %v", u)
	h.Respond(ctx, http.StatusOK, dto.ToUserResponse(u))
	return nil
}

func (h *UserProfileHandler) POST(ctx *routing.Context) error {
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

	u, err := h.userRepository.Update(&entities.User{
		Nickname: nickname,
		About:    req.About,
		Email:    req.Email,
		Fullname: req.Fullname,
	})

	if err != nil {
		h.UsecaseError(ctx, err, codesByErrorsPOST)
		return nil
	}

	h.Log(ctx).Debugf("update user %v", u)
	h.Respond(ctx, http.StatusOK, dto.ToUserResponse(u))
	return nil
}
