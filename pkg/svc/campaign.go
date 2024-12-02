package svc

import (
	"context"
	"seriouspoop/greedygame/go-common/utils"
	"seriouspoop/greedygame/pkg/model"
	"slices"

	"github.com/rs/zerolog"
)

func (s *Svc) GetActiveCampaignForDelivery(ctx context.Context, app, os, country string) ([]*model.Campaign, error) {
	log := s.logger.WithCtxLogger(ctx)
	if app == "" || os == "" || country == "" {
		log.Error().Err(ErrImportantFieldMissing).Msg("empty or invalid app, os or country")
		return nil, ErrImportantFieldMissing
	}

	targetingRules, err := s.db.GetTargetingRules(ctx)
	if err != nil {
		log.Error().Err(err).Msg("error while getting targeting rules")
		return nil, err
	}

	campaignIDs := []model.CampaignID{}

	for _, rule := range targetingRules {
		if s.matchInExclude(rule.Country, country) || s.matchInExclude(rule.OS, os) || s.matchInExclude(rule.App, app) {
			continue
		} else if s.matchInInclude(rule.Country, country) &&
			s.matchInInclude(rule.OS, os) &&
			s.matchInInclude(rule.App, app) {
			campaignIDs = append(campaignIDs, rule.CampaignID)
		}
	}

	if len(campaignIDs) == 0 {
		log.Debug().Msg("no campaignIDs found")
		return nil, ErrNoData
	}

	log.Debug().
		Dict("target", zerolog.Dict().
			Str("app", app).
			Str("country", country).
			Str("os", os)).
		Strs("campaign-ids", utils.StringSlice(campaignIDs)).
		Msg("campaign IDs found for the given target")

	campaigns, err := s.db.GetCampaignFromCIDs(ctx, campaignIDs, model.StatusActive)
	if err != nil {
		log.Error().Err(err).Strs("campaign-ids", utils.StringSlice(campaignIDs)).Msg("error while getting campaigns with cids")
		return nil, err
	}

	return campaigns, nil
}

func (s *Svc) matchInInclude(d *model.Dimensions, target string) bool {
	if d == nil {
		return true
	}
	return (len(d.Include) == 0 || slices.Contains(d.Include, target))
}

func (s *Svc) matchInExclude(d *model.Dimensions, target string) bool {
	if d == nil {
		return false
	} else if len(d.Exclude) == 0 {
		return false
	}
	return slices.Contains(d.Exclude, target)
}
