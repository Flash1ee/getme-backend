package handler_interfaces

import (
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

type FastHTTPFunc func(ctx *fasthttp.RequestCtx)

func (f FastHTTPFunc) ServeHTTP(ctx *routing.Context) error {
	f(ctx.RequestCtx)
	return nil
}

func (f FastHTTPFunc) ServeFastHTTP(ctx *fasthttp.RequestCtx) {
	f(ctx)
}

type HandlerFunc func(ctx *routing.Context) error

func (f HandlerFunc) ServeHTTP(ctx *routing.Context) error {
	return f(ctx)
}

type Handler interface {
	ServeHTTP(ctx *routing.Context) error
}
type HMiddlewareFunc func(Handler) Handler
type HFMiddlewareFunc func(HandlerFunc) HandlerFunc
