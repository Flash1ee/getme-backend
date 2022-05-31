package plans_task_handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	plans_usecase "getme-backend/internal/app/plans/usecase"
	session_client "getme-backend/internal/microservices/auth/delivery/grpc/client"
	middleware2 "getme-backend/internal/microservices/auth/sessions/middleware"
	"getme-backend/internal/pkg/adapter/echo_adapter"
	bh "getme-backend/internal/pkg/handler"
	"getme-backend/internal/pkg/handler/handler_errors"
)

type PlanIDTaskHandler struct {
	sessionClient session_client.AuthCheckerClient
	plansUsecase  plans_usecase.Usecase
	bh.BaseHandler
}

func NewPlanTaskHandler(log *logrus.Logger, sClient session_client.AuthCheckerClient, plansUs plans_usecase.Usecase) *PlanIDTaskHandler {
	h := &PlanIDTaskHandler{
		BaseHandler:   *bh.NewBaseHandler(log),
		sessionClient: sClient,
		plansUsecase:  plansUs,
	}

	h.AddMethod(http.MethodGet, h.GET, echo_adapter.WrapMiddlewareToFunc(middleware2.NewSessionMiddleware(sClient, log).CheckFunc))

	return h
}

func (h *PlanIDTaskHandler) GET(ctx echo.Context) error {
	userID, ok := ctx.Request().Context().Value("user_id").(int64)
	if !ok {
		h.Log(ctx.Request()).Error("can not get user_id from context")
		h.Error(ctx, http.StatusInternalServerError, handler_errors.InternalError)
		return handler_errors.InternalError
	}
	//req := &dto.RequestPlan{}
	//_, status := h.GetParamToStruct(ctx, req)
	//if status != bh.OK {
	//	h.Log(ctx.Request()).Warnf("can not parse query param")
	//	h.Error(ctx, http.StatusBadRequest, handler_errors.InvalidQueries)
	//	return handler_errors.InvalidQueries
	//}
	//if err := h.Validator.Struct(req); err != nil {
	//	h.Log(ctx.Request()).Warnf("can not validate query param, err = %v", err)
	//	h.Error(ctx, http.StatusBadRequest, handler_errors.InvalidQueries)
	//	return handler_errors.InvalidQueries
	//}
	planID, _, err := h.GetInt64FromParam(ctx, "id")
	if err != nil {
		h.Log(ctx.Request()).Warnf("PlanIDTaskHandler: can not parse path param ID %s", err)
		h.Error(ctx, http.StatusBadRequest, handler_errors.InvalidParameters)
		return handler_errors.InvalidParameters
	}

	res, err := h.plansUsecase.GetPlanWithTasks(userID, planID)
	if err != nil {
		h.UsecaseError(ctx, err, codeByErrGET)
		return err
	}
	//@TODO если ментор, один response/ менти - другой
	//if req.Role == dto.MentorAlias {
	//	h.Respond(ctx, http.StatusOK, dto.ToPlanWithSkillResponseByMentor(res))
	//} else {
	//	h.Respond(ctx, http.StatusOK, dto.ToPlanWithSkillsResponseByMentee(res))
	//}

	return nil
}
