package skills_info_handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"getme-backend/internal/app/skill/dto"
	"getme-backend/internal/app/skill/usecase"
	session_client "getme-backend/internal/microservices/auth/delivery/grpc/client"
	"getme-backend/internal/microservices/auth/sessions/middleware"
	"getme-backend/internal/pkg/adapter/echo_adapter"
	bh "getme-backend/internal/pkg/handler"
	"getme-backend/internal/pkg/handler/handler_errors"
)

type InfoHandler struct {
	sessionClient session_client.AuthCheckerClient
	skillUsecase  skill_usecase.Usecase
	bh.BaseHandler
}

func NewInfoHandler(log *logrus.Logger, sClient session_client.AuthCheckerClient, ucSkill skill_usecase.Usecase) *InfoHandler {
	h := &InfoHandler{
		BaseHandler:   *bh.NewBaseHandler(log),
		sessionClient: sClient,
		skillUsecase:  ucSkill,
	}
	h.AddMiddleware(echo_adapter.WrapMiddleware(middleware.NewSessionMiddleware(h.sessionClient, log).CheckFunc))

	h.AddMethod(http.MethodGet, h.GET)
	h.AddMethod(http.MethodPost, h.POST)

	return h
}

func (h *InfoHandler) GET(ctx echo.Context) error {
	skills, err := h.skillUsecase.GetAllSkills()
	if err != nil {
		h.Log(ctx.Request()).Warnf("Skills InfoHandler error method GET; error -  %s", err)
		h.Error(ctx, http.StatusUnprocessableEntity, handler_errors.InternalError)
		return handler_errors.InternalError
	}
	h.Log(ctx.Request()).Debugf("all skills : %v", skills)
	h.Respond(ctx, http.StatusOK, dto.ToSkillsResponseFromUsecase(skills))
	return nil
}
func (h *InfoHandler) POST(ctx echo.Context) error {
	userID := ctx.Request().Context().Value("user_id")
	if userID == nil {
		h.Log(ctx.Request()).Error("can not get user_id from context")
		h.Error(ctx, http.StatusInternalServerError, handler_errors.InternalError)
		return handler_errors.InternalError
	}
	req := &dto.RequestUsersBySkills{}
	_, status := h.GetParamToStruct(ctx, req)
	if status != bh.OK {
		h.Log(ctx.Request()).Errorf("SKILLS Handler : Error get params for users by skills request %v\n", req)
		ctx.Response().WriteHeader(http.StatusBadRequest)
		return nil
	}
	if len(req.Skills) == 0 {
		h.Log(ctx.Request()).Errorf("SKILLS Handler : No skills in query params %v\n", req.Skills)
		ctx.Response().WriteHeader(http.StatusBadRequest)
		return nil
	}

	skills, err := h.skillUsecase.GetAllSkills()
	if err != nil {
		h.Log(ctx.Request()).Warnf("Skills InfoHandler error method GET; error -  %s", err)
		h.Error(ctx, http.StatusUnprocessableEntity, handler_errors.InternalError)
		return handler_errors.InternalError
	}
	h.Log(ctx.Request()).Debugf("all skills : %v", skills)
	h.Respond(ctx, http.StatusOK, dto.ToSkillsResponseFromUsecase(skills))
	return nil
}
