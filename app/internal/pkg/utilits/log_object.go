package utilits

import (
	"net/http"

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

func (l *LogObject) Log(r *http.Request) *logrus.Entry {
	if r == nil {
		return l.log.WithField("type", "base_log")
	}
	ctxLogger := r.Context().Value("logger")
	logger := l.log.WithField("urls", r.URL)
	if ctxLogger != nil {
		if log, ok := ctxLogger.(*logrus.Entry); ok {
			logger = log
		}
	}
	return logger
}
