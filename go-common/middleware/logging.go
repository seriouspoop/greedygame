package middleware

import (
	"net/http"
	"seriouspoop/greedygame/go-common/logging"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Log struct {
	logger *logging.Logger
	level  zapcore.Level
}

func NewLog(logger *logging.Logger, level zapcore.Level) *Log {
	return &Log{logger: logger, level: level}
}

func (l *Log) LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		elapsed := time.Since(start)

		l.logger.Ctx(r.Context()).
			Log(l.level, "http request completed", zap.Int64("elapsed", elapsed.Nanoseconds()))
	})
}
