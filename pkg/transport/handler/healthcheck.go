package handler

import (
	"net/http"
	"seriouspoop/greedygame/pkg/svc"
)

func HealthCheck(s *svc.Svc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if s.IsUnhealthy(r.Context()) {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte("service unhealthy"))
		}
		w.Write([]byte("service healthy"))
	}
}
