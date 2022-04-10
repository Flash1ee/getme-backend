package delivery

import (
	"net/http"

	"github.com/pkg/errors"
	routing "github.com/qiangxue/fasthttp-routing"
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

func (h *ErrorConvertor) UsecaseError(ctx *routing.Context, usecaseErr error, codeByErr CodeMap) {
	var generalError *app.GeneralError
	//orginalError := usecaseErr
	if errors.As(usecaseErr, &generalError) {
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

	//h.Log(ctx).Logf(respond.Level, "Gotted error: %v", orginalError)
	h.Error(ctx, respond.Code, respond.Error)
}

func (h *ErrorConvertor) HandlerError(ctx *routing.Context, code int, err error) {
	//h.Log(ctx).Errorf("Gotted error: %v", err)

	var generalError *app.GeneralError
	if errors.As(err, &generalError) {
		err = errors.Cause(err).(*app.GeneralError).Err
	}
	h.Error(ctx, code, err)
}
