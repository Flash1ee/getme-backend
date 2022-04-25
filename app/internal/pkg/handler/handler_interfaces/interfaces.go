package handler_interfaces

import (
	"github.com/labstack/echo/v4"
)

type Handler interface {
	ServeHTTP(ctx echo.Context) error
}

//type HandlerFunc func(http.ResponseWriter, *http.Request)

type HandlerFunc func(ctx echo.Context) error

func (f HandlerFunc) ServeHTTP(ctx echo.Context) error {
	return f(ctx)
}

type HMiddlewareFunc func(handler Handler) Handler

//type HFMiddlewareFunc func(HandlerFunc) HandlerFunc
type HFMiddlewareFunc echo.MiddlewareFunc

//func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
//	return f(w, r)
//}
//

//type HMiddlewareFunc func(Handler) Handler
//type HFMiddlewareFunc func(HandlerFunc) HandlerFunc
