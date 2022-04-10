package app

import (
	routing "github.com/qiangxue/fasthttp-routing"
)

type Handler interface {
	ServeHTTP(ctx *routing.Context) error
	Connect(router *routing.Route)
}
