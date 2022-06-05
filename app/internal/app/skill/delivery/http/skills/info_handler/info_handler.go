package skills_info_handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"getme-backend/internal/app/skill/dto"
	"getme-backend/internal/app/skill/usecase"
	session_client "getme-backend/internal/microservices/auth/delivery/grpc/client"
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

	h.AddMethod(http.MethodGet, h.GET)

	return h
}

func (h *InfoHandler) GET(ctx echo.Context) error {
	skills, err := h.skillUsecase.GetAllSkills()
	if err != nil {
		h.Log(ctx.Request()).Warnf("Skills InfoHandler error method GET; error -  %s", err)
		h.Error(ctx, http.StatusInternalServerError, handler_errors.InternalError)
		return handler_errors.InternalError
	}
	h.Log(ctx.Request()).Debugf("all skills : %v", skills)
	h.Respond(ctx, http.StatusOK, dto.ToSkillsResponseFromUsecase(skills))
	return nil
}
