package task_id_handler

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

type TaskIdHandler struct {
	sessionClient session_client.AuthCheckerClient
	taskUsecase   task_usecase.Usecase
	bh.BaseHandler
}

func NewTaskIdHandler(log *logrus.Logger, sClient session_client.AuthCheckerClient, taskUs task_usecase.Usecase) *TaskIdHandler {
	h := &TaskIdHandler{
		BaseHandler:   *bh.NewBaseHandler(log),
		sessionClient: sClient,
		taskUsecase:   taskUs,
	}

	h.AddMethod(http.MethodPut, h.PUT, echo_adapter.WrapMiddlewareToFunc(middleware2.NewSessionMiddleware(sClient, log).CheckFunc))

	return h
}

func (h *TaskIdHandler) PUT(ctx echo.Context) error {
	userID, ok := ctx.Request().Context().Value("user_id").(int64)
	if !ok {
		h.Log(ctx.Request()).Error("can not get user_id from context")
		h.Error(ctx, http.StatusInternalServerError, handler_errors.InternalError)
		return handler_errors.InternalError
	}

	taskId, status, err := h.GetInt64FromParam(ctx, "id")
	if err != nil {
		h.Error(ctx, status, err)
		return err
	}

	if len(ctx.ParamValues()) > 1 {
		h.Log(ctx.Request()).Warnf("Too many parametres %v", ctx.ParamValues())
		h.Error(ctx, http.StatusBadRequest, handler_errors.InvalidParameters)
		return handler_errors.InvalidParameters
	}

	req := &dto.TaskUsecaseDTO{}

	req.ID = taskId

	err = h.taskUsecase.ApplyTask(userID, *req)
	if err != nil {
		h.UsecaseError(ctx, err, codeByErrPUT)
		return err
	}

	h.Respond(ctx, http.StatusOK, dto.TaskIDResponse{ID: taskId})
	return nil
}
