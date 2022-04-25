package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/labstack/echo/v4"

	hf "getme-backend/internal/pkg/handler/handler_interfaces"
)

type InterfaceBaseHandler interface {
	AddMiddleware(middleware ...hf.HMiddlewareFunc)
	AddMethod(method string, handlerMethod hf.HandlerFunc, middlewares ...hf.HFMiddlewareFunc)
	applyHFMiddleware(handlerMethod hf.HandlerFunc, middlewares ...hf.HFMiddlewareFunc) hf.HandlerFunc
	applyMiddleware(handler hf.Handler) hf.Handler
	getListMethods() []string
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	Connect(route *mux.Route)
	add(path string, handler hf.HandlerFunc, route *echo.Group)
}
