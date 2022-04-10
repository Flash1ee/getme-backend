package middleware

import (
	"net/http"
	"runtime/debug"
	"time"

	"github.com/google/uuid"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/sirupsen/logrus"

	hf "getme-backend/internal/pkg/handler/handler_interfaces"
	"getme-backend/internal/pkg/utilits"
)

type UtilitiesMiddleware struct {
	log utilits.LogObject
}

func NewUtilitiesMiddleware(log *logrus.Logger) UtilitiesMiddleware {
	return UtilitiesMiddleware{
		log: utilits.NewLogObject(log),
	}
}

func (mw *UtilitiesMiddleware) CheckPanic() hf.Handler {
	return hf.HandlerFunc(func(ctx *routing.Context) error {
		defer func(log *logrus.Entry, ctx *routing.Context) {
			if err := recover(); err != nil {
				responseErr := http.StatusInternalServerError

				log.Errorf("detacted critical error: %v, with stack: %s", err, debug.Stack())
				ctx.SetStatusCode(responseErr)
			}
		}(mw.log.Log(ctx), ctx)
		return ctx.Next()
	})
}

func (mw *UtilitiesMiddleware) UpgradeLogger() hf.Handler {
	return hf.HandlerFunc(func(ctx *routing.Context) error {
		start := time.Now()
		upgradeLogger := mw.log.BaseLog().WithFields(logrus.Fields{
			"urls":        ctx.URI().String(),
			"method":      string(ctx.Method()),
			"remote_addr": ctx.RemoteAddr(),
			"work_time":   time.Since(start).Milliseconds(),
			"req_id":      uuid.NewUUID(),
		})
		ctx.SetUserValue("logger", upgradeLogger)
		upgradeLogger.Info("Log was upgraded")

		err := ctx.Next()

		executeTime := time.Since(start).Milliseconds()
		upgradeLogger.Infof("work time [ms]: %v", executeTime)
		return err
	})
}
