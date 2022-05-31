package task_handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"getme-backend/internal/app/task/dto"
	task_usecase "getme-backend/internal/app/task/usecase"
	session_client "getme-backend/internal/microservices/auth/delivery/grpc/client"
	middleware2 "getme-backend/internal/microservices/auth/sessions/middleware"
	"getme-backend/internal/pkg/adapter/echo_adapter"
	bh "getme-backend/internal/pkg/handler"
	"getme-backend/internal/pkg/handler/handler_errors"
)

type TaskHandler struct {
	sessionClient session_client.AuthCheckerClient
	taskUsecase   task_usecase.Usecase
	bh.BaseHandler
}

func NewTaskHandler(log *logrus.Logger, sClient session_client.AuthCheckerClient, taskUs task_usecase.Usecase) *TaskHandler {
	h := &TaskHandler{
		BaseHandler:   *bh.NewBaseHandler(log),
		sessionClient: sClient,
		taskUsecase:   taskUs,
	}

	h.AddMethod(http.MethodPost, h.POST, echo_adapter.WrapMiddlewareToFunc(middleware2.NewSessionMiddleware(sClient, log).CheckFunc))
	h.AddMethod(http.MethodGet, h.GET, echo_adapter.WrapMiddlewareToFunc(middleware2.NewSessionMiddleware(sClient, log).CheckFunc))

	return h
}

func (h *TaskHandler) POST(ctx echo.Context) error {
	userID, ok := ctx.Request().Context().Value("user_id").(int64)
	if !ok {
		h.Log(ctx.Request()).Error("can not get user_id from context")
		h.Error(ctx, http.StatusInternalServerError, handler_errors.InternalError)
		return handler_errors.InternalError
	}
	req := &dto.RequestCreateTask{}
	_, status := h.GetParamToStruct(ctx, req)
	if status != bh.OK {
		h.Log(ctx.Request()).Warnf("can not parse query param")
		h.Error(ctx, http.StatusBadRequest, handler_errors.InvalidQueries)
		return handler_errors.InvalidQueries
	}
	if err := h.Validator.Struct(req); err != nil {
		h.Log(ctx.Request()).Warnf("can not validate body param, err = %v", err)
		h.Error(ctx, http.StatusBadRequest, handler_errors.InvalidQueries)
		return handler_errors.InvalidQueries
	}

	res, err := h.taskUsecase.Create(userID, *req.ToRequestCreateTaskDTO())
	if err != nil {
		h.UsecaseError(ctx, err, codeByErrPOST)
		return err
	}

	h.Respond(ctx, http.StatusCreated, dto.TaskIDResponse{ID: res})
	return nil
}

func (h *TaskHandler) GET(ctx echo.Context) error {
	userID, ok := ctx.Request().Context().Value("user_id").(int64)
	if !ok {
		h.Log(ctx.Request()).Error("can not get user_id from context")
		h.Error(ctx, http.StatusInternalServerError, handler_errors.InternalError)
		return handler_errors.InternalError
	}
	req := &dto.RequestCreateTask{}
	_, status := h.GetParamToStruct(ctx, req)
	if status != bh.OK {
		h.Log(ctx.Request()).Warnf("can not parse query param")
		h.Error(ctx, http.StatusBadRequest, handler_errors.InvalidQueries)
		return handler_errors.InvalidQueries
	}
	if err := h.Validator.Struct(req); err != nil {
		h.Log(ctx.Request()).Warnf("can not validate body param, err = %v", err)
		h.Error(ctx, http.StatusBadRequest, handler_errors.InvalidQueries)
		return handler_errors.InvalidQueries
	}

	res, err := h.taskUsecase.Create(userID, *req.ToRequestCreateTaskDTO())
	if err != nil {
		h.UsecaseError(ctx, err, codeByErrPOST)
		return err
	}

	h.Respond(ctx, http.StatusCreated, dto.TaskIDResponse{ID: res})
	return nil
}
