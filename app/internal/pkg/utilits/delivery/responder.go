package delivery

import (
	"github.com/labstack/echo/v4"
	"github.com/mailru/easyjson"

	"getme-backend/internal/pkg/utilits"
)

//go:generate easyjson -disallow_unknown_fields responder.go

//easyjson:json
type ErrResponse struct {
	Err string `json:"message"`
}

type Responder struct {
	utilits.LogObject
}

func (h *Responder) Error(ctx echo.Context, code int, err error) {
	h.Respond(ctx, code, ErrResponse{Err: err.Error()})
}

func (h *Responder) Respond(ctx echo.Context, code int, data easyjson.Marshaler) {
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	ctx.Response().WriteHeader(code)
	if data != nil {
		_, _, err := easyjson.MarshalToHTTPResponseWriter(data, ctx.Response())
		if err != nil {
			h.Log(ctx.Request()).Error(err)
		}
	}
	logUser, _ := easyjson.Marshal(data)
	h.Log(ctx.Request()).Info("Respond data: ", string(logUser))
}
