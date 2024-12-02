package middleware

import (
	"net/http"
	"seriouspoop/greedygame/go-common/logging"
	"time"

	"github.com/rs/zerolog"
)

type Log struct {
	logger *logging.Logger
	level  zerolog.Level
}

func NewLog(logger *logging.Logger, level zerolog.Level) *Log {
	return &Log{logger: logger, level: level}
}

func (l *Log) LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		elapsed := time.Since(start)

		l.logger.WithCtxLogger(r.Context()).
			WithLevel(l.level).
			Int64("elapsed", elapsed.Nanoseconds()).Send()
	})
}
