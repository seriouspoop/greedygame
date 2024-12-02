package handler

import (
	"net/http"
	"seriouspoop/greedygame/go-common/logging"
	"seriouspoop/greedygame/go-common/middleware"
	"seriouspoop/greedygame/pkg/svc"
)

type DeliveryResponse struct {
	CampaignID string `json:"cid"`
	Image      string `json:"img"`
	CTA        string `json:"cta"`
}

func Delivery(s servicer, logger *logging.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := logger.WithCtxLogger(ctx)
		app := r.URL.Query().Get("app")
		os := r.URL.Query().Get("os")
		country := r.URL.Query().Get("country")

		if app == "" || os == "" || country == "" {
			writeErrorResponse(svc.ErrImportantFieldMissing, r, w)
			return
		}

		campaigns, err := s.GetActiveCampaignForDelivery(ctx, app, os, country)

		if err != nil {
			logger.Error().Err(err).Msg("error while getting campaigns for delivery")
			writeErrorResponse(err, r, w)
			return
		}

		response := []DeliveryResponse{}
		for _, campaign := range campaigns {
			delivery := DeliveryResponse{
				CampaignID: campaign.ID.String(),
				Image:      campaign.Image.String(),
				CTA:        campaign.CTA,
			}
			response = append(response, delivery)
		}

		middleware.WriteJsonHttpResponse(ctx, w, http.StatusOK, response)
	}
}
