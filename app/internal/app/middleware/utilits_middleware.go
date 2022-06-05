package middleware

import (
	"context"
	"net/http"
	"runtime/debug"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"

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

func (mw UtilitiesMiddleware) CheckPanic(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func(log *logrus.Entry, w http.ResponseWriter) {
			if err := recover(); err != nil {
				responseErr := http.StatusInternalServerError
				log.Errorf("detacted critical error: %v, with stack: %s", err, debug.Stack())
				w.WriteHeader(responseErr)
			}
		}(mw.log.Log(r), w)
		h.ServeHTTP(w, r)
	})
}

//func (mw *UtilitiesMiddleware) UpgradeLogger(hf hf.HandlerFunc) hf.HandlerFunc {
//	return func(ctx echo.Context) error {
//		start := time.Now()
//		upgradeLogger := mw.log.BaseLog().WithFields(logrus.Fields{
//			"urls":        ctx.Request().RequestURI,
//			"method":      ctx.Request().Method,
//			"remote_addr": ctx.Request().RemoteAddr,
//			"work_time":   time.Since(start).Milliseconds(),
//			"req_id":      uuid.NewUUID(),
//		})
//		ctx.Set("logger", upgradeLogger)
//		upgradeLogger.Info("Log was upgraded")
//
//		err := hf
//
//		executeTime := time.Since(start).Milliseconds()
//		upgradeLogger.Infof("work time [ms]: %v", executeTime)
//		return err
//	})
//}
func (mw *UtilitiesMiddleware) UpgradeLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		upgradeLogger := mw.log.BaseLog().WithFields(logrus.Fields{
			"urls":        r.URL,
			"method":      r.Method,
			"remote_addr": r.RemoteAddr,
			"work_time":   time.Since(start).Milliseconds(),
			"req_id":      uuid.NewV4(),
		})

		r = r.WithContext(context.WithValue(r.Context(), "logger", upgradeLogger))
		upgradeLogger.Info("Log was upgraded")

		handler.ServeHTTP(w, r)

		executeTime := time.Since(start).Milliseconds()
		upgradeLogger.Infof("work time [ms]: %v", executeTime)

	})
}
