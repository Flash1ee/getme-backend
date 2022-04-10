package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/mailru/easyjson"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"

	"getme-backend/internal/app"
	"getme-backend/internal/pkg/handler/handler_errors"
	"getme-backend/internal/pkg/utilits/delivery"
)

type Pagination struct {
	Limit int64
	Desc  bool
	Since string
}

const (
	EmptyQuery   = -2
	DefaultLimit = 100
)

type HelpHandlers struct {
	delivery.ErrorConvertor
}

func (h *HelpHandlers) PrintRequest(ctx *routing.Context) {
	h.Log(ctx).Infof("Request: %s. From URL: %s", ctx.Method(), string(ctx.URI().Host())+string(ctx.Path()))
}

// GetInt64FromParam HTTPErrors
//		Status 400 handler_errors.InvalidParameters
func (h *HelpHandlers) GetInt64FromParam(ctx *routing.Context, name string) (int64, int, error) {
	number := ctx.Param(name)
	numberInt, err := strconv.ParseInt(number, 10, 64)
	if number == "" || err != nil {
		//h.Log(ctx).Infof("can't get parametrs %s, was got %v)", name, number)
		return app.InvalidInt, http.StatusBadRequest, handler_errors.InvalidParameters
	}
	return numberInt, app.InvalidInt, nil
}

// GetPaginationFromQuery Expected api param:
// 	Default value for limit - 100
//	Param since query any false "start number of values"
// 	Param limit query uint64 false "number values to return"
//	Param desc  query bool false "
// Errors:
// 	Status 400 handler_errors.InvalidQueries
func (h *HelpHandlers) GetPaginationFromQuery(ctx *routing.Context) (*Pagination, int, error) {
	limit, code, err := h.GetInt64FromQueries(ctx, "limit")
	if err != nil {
		return nil, code, err
	}

	if limit == EmptyQuery {
		limit = DefaultLimit
	}

	desc := h.GetBoolFromQueries(ctx, "desc")

	since, info := h.GetStringFromQueries(ctx, "since")
	if info == EmptyQuery {
		since = ""
	}
	return &Pagination{Since: since, Desc: desc, Limit: limit}, app.InvalidInt, nil
}

// GetInt64FromQueries HTTPErrors
//		Status 400 handler_errors.InvalidQueries
func (h *HelpHandlers) GetInt64FromQueries(ctx *routing.Context, name string) (int64, int, error) {
	number, err := ctx.URI().QueryArgs().GetUint(name)
	if err != nil {
		if err == fasthttp.ErrNoArgValue {
			return EmptyQuery, app.InvalidInt, nil
		}
		return app.InvalidInt, http.StatusBadRequest, handler_errors.InvalidQueries
	}

	return int64(number), app.InvalidInt, nil
}

// GetBoolFromQueries HTTPErrors
//		Status 400 handler_errors.InvalidQueries
func (h *HelpHandlers) GetBoolFromQueries(ctx *routing.Context, name string) bool {
	res := ctx.URI().QueryArgs().GetBool(name)
	return res
}

// GetStringFromQueries HTTPErrors
//		Status 400 handler_errors.InvalidQueries
func (h *HelpHandlers) GetStringFromQueries(ctx *routing.Context, name string) (string, int) {
	value := ctx.URI().QueryArgs().Peek(name)
	if value == nil {
		return "", EmptyQuery
	}

	return string(value), app.InvalidInt
}

// GetStringFromParam HTTPErrors
//		Status 400 handler_errors.InvalidQueries
func (h *HelpHandlers) GetStringFromParam(ctx *routing.Context, name string) (string, int) {
	value := ctx.Param(name)
	if value == "" {
		return "", EmptyQuery
	}

	return value, app.InvalidInt
}

// GetArrayStringFromQueries HTTPErrors
//		Status 400 handler_errors.InvalidQueries
func (h *HelpHandlers) GetArrayStringFromQueries(ctx *routing.Context, name string) ([]string, int) {
	values := ctx.URI().QueryArgs().PeekMultiBytes([]byte(name))
	if values == nil {
		return nil, EmptyQuery
	}

	var res []string
	for _, value := range values {
		res = append(res, strings.Split(string(value), ",")...)
	}
	return res, app.InvalidInt
}

func (h *HelpHandlers) GetRequestBody(ctx *routing.Context, reqStruct easyjson.Unmarshaler) error {
	return easyjson.Unmarshal(ctx.PostBody(), reqStruct)
}
