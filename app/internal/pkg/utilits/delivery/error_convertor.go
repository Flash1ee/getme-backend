package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"getme-backend/internal/app"
)

type RespondError struct {
	Code  int
	Error error
	Level logrus.Level
}

type CodeMap map[error]RespondError

type ErrorConvertor struct {
	Responder
}

func (h *ErrorConvertor) UsecaseError(ctx echo.Context, usecaseErr error, codeByErr CodeMap) {
	var generalError *app.GeneralError
	//orginalError := usecaseErr
	if errors.As(usecaseErr, &generalError) {
		h.Log(ctx.Request()).Warnf("external error =  %v", errors.Cause(usecaseErr).(*app.GeneralError).ExternalErr)
		usecaseErr = errors.Cause(usecaseErr).(*app.GeneralError).Err
	}

	respond := RespondError{http.StatusServiceUnavailable,
		errors.New("UnknownError"), logrus.ErrorLevel}

	for err, respondErr := range codeByErr {
		if errors.Is(usecaseErr, err) {
			respond = respondErr
			break
		}
	}

	h.Log(ctx.Request()).Logf(respond.Level, "Gotted error: %v", usecaseErr)
	h.Error(ctx, respond.Code, respond.Error)
}

func (h *ErrorConvertor) HandlerError(ctx echo.Context, code int, err error) {
	h.Log(ctx.Request()).Errorf("Gotted error: %v", err)

	var generalError *app.GeneralError
	if errors.As(err, &generalError) {
		err = errors.Cause(err).(*app.GeneralError).Err
	}
	h.Error(ctx, code, err)
}
