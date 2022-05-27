package skills_user_handler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"getme-backend/internal/app/skill/dto"
	skill_usecase "getme-backend/internal/app/skill/usecase"
	dto2 "getme-backend/internal/app/user/dto"
	session_client "getme-backend/internal/microservices/auth/delivery/grpc/client"
	bh "getme-backend/internal/pkg/handler"
)

type UserHandler struct {
	sessionClient session_client.AuthCheckerClient
	skillUsecase  skill_usecase.Usecase
	bh.BaseHandler
}

func NewUserHandler(log *logrus.Logger, sClient session_client.AuthCheckerClient, ucSkill skill_usecase.Usecase) *UserHandler {
	h := &UserHandler{
		BaseHandler:   *bh.NewBaseHandler(log),
		sessionClient: sClient,
		skillUsecase:  ucSkill,
	}
	//h.AddMiddleware(echo_adapter.WrapMiddleware(middleware.NewSessionMiddleware(h.sessionClient, log).CheckFunc))

	h.AddMethod(http.MethodGet, h.GET)

	return h
}

func (h *UserHandler) GET(ctx echo.Context) error {
	req := &dto.RequestUsersBySkills{}
	_, status := h.GetParamToStruct(ctx, req)
	if status != bh.OK {
		h.Log(ctx.Request()).Errorf("skills_user_handler GET : Error get params for users by skills request %v\n", req)
		ctx.Response().WriteHeader(http.StatusBadRequest)
		return nil
	}

	//if len(req.Skills) == 0 {
	//	h.Log(ctx.Request()).Errorf("SKILLS Handler : No skills in query params %v\n", req.Skills)
	//	ctx.Response().WriteHeader(http.StatusBadRequest)
	//	return nil
	//}
	if len(req.Skills) == 1 {
		req.Skills = strings.Split(req.Skills[0], ",")
	}

	users, err := h.skillUsecase.GetUsersBySkills(req.ToSkillUsecase())
	if err != nil {
		h.Log(ctx.Request()).Warnf("skills_user_handler error method GET; error -  %s", err)
		h.UsecaseError(ctx, err, codeByErrorGET)
		return err
	}
	h.Log(ctx.Request()).Debugf("users: %v\n skills : %v\n", users, req.Skills)
	h.Respond(ctx, http.StatusOK, dto2.ToUsersWithSkillResponse(users))
	return nil
}
