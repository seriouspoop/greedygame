package middleware

import (
	"net/http"
	"time"

	"github.com/rs/zerolog"
)

type Log struct {
	logger *zerolog.Logger
	level  zerolog.Level
}

func NewLog(logger *zerolog.Logger, level zerolog.Level) *Log {
	return &Log{logger: logger, level: level}
}

func (l *Log) LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		elapsed := time.Since(start)

		l.logger.WithLevel(l.level).
			Ctx(r.Context()).
			Int64("elapsed", elapsed.Nanoseconds()).Send()
	})
}
