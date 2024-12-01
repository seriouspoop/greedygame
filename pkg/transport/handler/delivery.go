package handler

import (
	"net/http"
	"seriouspoop/greedygame/go-common/logging"
	"seriouspoop/greedygame/pkg/svc"
)

func Delivery(s servicer, logger *logging.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// logger := logger.With().Ctx(ctx).Logger()
		app := r.URL.Query().Get("app")
		os := r.URL.Query().Get("os")
		country := r.URL.Query().Get("country")

		if app == "" || os == "" || country == "" {
			writeErrorResponse(svc.ErrImportantFieldMissing, r, w)
			return
		}

		_, _ = s.GetCampaignForDelivery(ctx, app, os, country)
	}
}
