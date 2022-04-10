package delivery

import (
	"net/http"
	"strconv"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jwriter"
	routing "github.com/qiangxue/fasthttp-routing"

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

func (h *Responder) Error(ctx *routing.Context, code int, err error) {
	h.Respond(ctx, code, ErrResponse{Err: err.Error()})
}

func (h *Responder) Respond(ctx *routing.Context, code int, data easyjson.Marshaler) {
	wasErr := false
	if data != nil {
		jw := jwriter.Writer{}
		data.MarshalEasyJSON(&jw)
		if jw.Error != nil {
			//h.Log(ctx).Error(jw.Error)
			wasErr = true
		} else {
			ctx.Response.Header.Set("Content-Type", "application/json")
			ctx.Response.Header.Set("Content-Length", strconv.Itoa(jw.Size()))

			_, err := jw.DumpTo(ctx.Response.BodyWriter())
			if err != nil {
				//h.Log(ctx).Error(jw.Error)
				wasErr = true
			}
		}
	}
	if wasErr {
		ctx.SetStatusCode(http.StatusInternalServerError)
		return
	}
	ctx.SetStatusCode(code)
	logUser, _ := easyjson.Marshal(data)
	h.Log(ctx).Info("Respond data: ", string(logUser))
}
