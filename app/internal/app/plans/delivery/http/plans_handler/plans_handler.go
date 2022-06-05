package plans_handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	offer_usecase "getme-backend/internal/app/offer/usecase"
	"getme-backend/internal/app/plans/dto"
	plans_usecase "getme-backend/internal/app/plans/usecase"
	session_client "getme-backend/internal/microservices/auth/delivery/grpc/client"
	middleware2 "getme-backend/internal/microservices/auth/sessions/middleware"
	"getme-backend/internal/pkg/adapter/echo_adapter"
	bh "getme-backend/internal/pkg/handler"
	"getme-backend/internal/pkg/handler/handler_errors"
)

type PlanHandler struct {
	sessionClient session_client.AuthCheckerClient
	offerUsecase  offer_usecase.Usecase
	plansUsecase  plans_usecase.Usecase
	bh.BaseHandler
}

func NewPlanHandler(log *logrus.Logger, sClient session_client.AuthCheckerClient, offerUs offer_usecase.Usecase, plansUs plans_usecase.Usecase) *PlanHandler {
	h := &PlanHandler{
		BaseHandler:   *bh.NewBaseHandler(log),
		sessionClient: sClient,
		offerUsecase:  offerUs,
		plansUsecase:  plansUs,
	}

	h.AddMethod(http.MethodGet, h.GET, echo_adapter.WrapMiddlewareToFunc(middleware2.NewSessionMiddleware(sClient, log).CheckFunc))

	return h
}

func (h *PlanHandler) GET(ctx echo.Context) error {
	userID, ok := ctx.Request().Context().Value("user_id").(int64)
	if !ok {
		h.Log(ctx.Request()).Error("can not get user_id from context")
		h.Error(ctx, http.StatusInternalServerError, handler_errors.InternalError)
		return handler_errors.InternalError
	}
	req := &dto.RequestPlan{}
	_, status := h.GetParamToStruct(ctx, req)
	if status != bh.OK {
		h.Log(ctx.Request()).Warnf("can not parse query param")
		h.Error(ctx, http.StatusBadRequest, handler_errors.InvalidQueries)
		return handler_errors.InvalidQueries
	}
	if err := h.Validator.Struct(req); err != nil {
		h.Log(ctx.Request()).Warnf("can not validate query param, err = %v", err)
		h.Error(ctx, http.StatusBadRequest, handler_errors.InvalidQueries)
		return handler_errors.InvalidQueries
	}

	res, err := h.plansUsecase.GetPlansByRole(userID, req.Role)
	if err != nil {
		h.UsecaseError(ctx, err, codeByErrGET)
		return err
	}
	if req.Role == dto.MentorAlias {
		h.Respond(ctx, http.StatusOK, dto.ToPlanWithSkillResponseByMentor(res))
	} else {
		h.Respond(ctx, http.StatusOK, dto.ToPlanWithSkillsResponseByMentee(res))
	}

	return nil
}
