package app

import (
	"github.com/labstack/echo/v4"
)

type Handler interface {
	ServeHTTP(ctx echo.Context) error
	Connect(route *echo.Group, path string)
}
