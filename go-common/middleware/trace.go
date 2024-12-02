package middleware

import (
	"context"
	"net/http"
	"seriouspoop/greedygame/go-common/globals"

	"github.com/google/uuid"
)

func TraceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceId := r.Header.Get(globals.TraceIDHeaderKey.String())
		if traceId == "" {
			traceId = uuid.NewString()
		}
		r.Header.Set(globals.TraceIDHeaderKey.String(), traceId)
		ctx := r.Context()
		ctx = context.WithValue(ctx, globals.TraceIDContextKey, traceId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
