package handler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

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
		resultHandlerMethod = hf.HandlerFunc(middlewares[index](echo.HandlerFunc(resultHandlerMethod)))
	}
	return resultHandlerMethod
}

func (h *BaseHandler) applyMiddleware(handler hf.Handler) hf.HandlerFunc {
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

func (h *BaseHandler) add(path string, echoHandlerFunc echo.HandlerFunc, route *echo.Group) {
	wrapped := echoHandlerFunc

	for key := range h.handlerMethods {
		switch key {
		case GET:
			route.GET(path, wrapped)
			break
		case POST:
			route.POST(path, wrapped)
			break
		case PUT:
			route.PUT(path, wrapped)
			break
		case DELETE:
			route.DELETE(path, wrapped)
			break
		case OPTIONS:
			route.OPTIONS(path, wrapped)
			break
		}
	}
}

func (h *BaseHandler) Connect(route *echo.Group, path string) {
	h.add(path, echo.HandlerFunc(h.applyMiddleware(h)), route)
}

func (h *BaseHandler) ServeHTTP(ctx echo.Context) error {
	h.PrintRequest(ctx.Request())
	ok := true
	var handler hf.HandlerFunc
	handler, ok = h.handlerMethods[ctx.Request().Method]
	if ok {
		return handler(ctx)
	}
	h.Log(ctx.Request()).Errorf("Unexpected http method: %s", ctx.Request().Method)
	ctx.Response().Header().Set("Allow", strings.Join(h.getListMethods(), ", "))
	ctx.Response().WriteHeader(http.StatusInternalServerError)
	return echo.ErrInternalServerError
}
