package middleware

// func TraceMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		traceId := r.Header.Get(globals.HeaderKeyTraceID.String())
// 		if traceId == "" {
// 			traceId = uuid.NewString()
// 		}
// 		r.Header.Set(globals.HeaderKeyTraceID.String(), traceId)
// 		ctx := r.Context()
// 		ctx = context.WithValue(ctx, globals.ContextKeyTraceID, traceId)
// 		r = r.WithContext(ctx)
// 		next.ServeHTTP(w, r)
// 	})
// }
