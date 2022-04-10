package handler

import (
	"net/http"
	"strings"

	"github.com/qiangxue/fasthttp-routing"

	hf "getme-backend/internal/pkg/handler/handler_interfaces"
	"getme-backend/internal/pkg/utilits"
	"getme-backend/internal/pkg/utilits/delivery"

	"github.com/sirupsen/logrus"
)

const (
	GET     = http.MethodGet
	POST    = http.MethodPost
	PUT     = http.MethodPut
	DELETE  = http.MethodDelete
	OPTIONS = http.MethodOptions
)

type BaseHandler struct {
	handlerMethods map[string]hf.HandlerFunc
	middlewares    []hf.HMiddlewareFunc
	HelpHandlers
}

func NewBaseHandler(log *logrus.Logger) *BaseHandler {
	h := &BaseHandler{handlerMethods: map[string]hf.HandlerFunc{}, middlewares: []hf.HMiddlewareFunc{},
		HelpHandlers: HelpHandlers{
			ErrorConvertor: delivery.ErrorConvertor{
				Responder: delivery.Responder{
					LogObject: utilits.NewLogObject(log),
				},
			},
		},
	}
	return h
}

func (h *BaseHandler) AddMiddleware(middleware ...hf.HMiddlewareFunc) {
	h.middlewares = append(h.middlewares, middleware...)
}

func (h *BaseHandler) AddMethod(method string, handlerMethod hf.HandlerFunc, middlewares ...hf.HFMiddlewareFunc) {
	h.handlerMethods[method] = h.applyHFMiddleware(handlerMethod, middlewares...)
}

func (h *BaseHandler) applyHFMiddleware(handlerMethod hf.HandlerFunc,
	middlewares ...hf.HFMiddlewareFunc) hf.HandlerFunc {
	resultHandlerMethod := handlerMethod
	for index := len(middlewares) - 1; index >= 0; index-- {
		resultHandlerMethod = middlewares[index](resultHandlerMethod)
	}
	return resultHandlerMethod
}

func (h *BaseHandler) applyMiddleware(handler hf.Handler) routing.Handler {
	resultHandler := handler
	for index := len(h.middlewares) - 1; index >= 0; index-- {
		resultHandler = h.middlewares[index](resultHandler)
	}
	return resultHandler.ServeHTTP
}

func (h *BaseHandler) getListMethods() []string {
	var useMethods []string
	for key := range h.handlerMethods {
		useMethods = append(useMethods, key)
	}
	useMethods = append(useMethods, http.MethodOptions)
	return useMethods
}

func (h *BaseHandler) add(handler routing.Handler, route *routing.Route) {
	for key := range h.handlerMethods {
		switch key {
		case GET:
			route.Get(handler)
			break
		case POST:
			route.Post(handler)
			break
		case PUT:
			route.Put(handler)
			break
		case DELETE:
			route.Delete(handler)
			break
		case OPTIONS:
			route.Options(handler)
			break
		}
	}
}

func (h *BaseHandler) Connect(route *routing.Route) {
	h.add(h.applyMiddleware(h), route)
}

func (h *BaseHandler) ServeHTTP(ctx *routing.Context) error {
	h.PrintRequest(ctx)
	ok := true
	var hndlr hf.HandlerFunc

	hndlr, ok = h.handlerMethods[string(ctx.Method())]
	if ok {
		hndlr(ctx)
	} else {
		h.Log(ctx).Errorf("Unexpected http method: %s", ctx.Method())
		ctx.Response.Header.Set("Allow", strings.Join(h.getListMethods(), ", "))
		ctx.SetStatusCode(http.StatusInternalServerError)
	}
	return nil
}
