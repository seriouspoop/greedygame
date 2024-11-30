package svc

import (
	"context"
	"seriouspoop/greedygame/pkg/model"
)

func (s *Svc) GetCampaignForDelivery(ctx context.Context, app, os, country string) ([]*model.Campaign, error) {
	log := s.logger.With().Ctx(ctx).Logger()
	if app == "" || os == "" || country == "" {
		log.Error().Err(ErrImportantFieldMissing).Msg("empty or invalid app, os or country")
		return nil, ErrImportantFieldMissing
	}

	return nil, nil
}
