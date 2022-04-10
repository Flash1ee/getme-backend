package server

import (
	"getme-backend/internal/app"
)

type HandlerFactory interface {
	GetHandleUrls() *map[string]app.Handler
}
