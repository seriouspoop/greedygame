package handler

import (
	"net/http"
	"seriouspoop/greedygame/pkg/svc"

	"github.com/rs/zerolog"
)

func Delivery(s *svc.Svc, logger *zerolog.Logger) http.HandlerFunc {
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
