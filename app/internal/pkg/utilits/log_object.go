package utilits

import (
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/sirupsen/logrus"
)

type LogObject struct {
	log *logrus.Logger
}

func NewLogObject(log *logrus.Logger) LogObject {
	return LogObject{log: log}
}

func (l *LogObject) BaseLog() *logrus.Logger {
	return l.log
}

func (l *LogObject) Log(ctx *routing.Context) *logrus.Entry {
	if ctx == nil {
		return l.log.WithField("type", "base_log")
	}
	ctxLogger := ctx.UserValue("logger")
	logger := l.log.WithField("urls", ctx.URI())
	if ctxLogger != nil {
		if log, ok := ctxLogger.(*logrus.Entry); ok {
			logger = log
		}
	}
	return logger
}
