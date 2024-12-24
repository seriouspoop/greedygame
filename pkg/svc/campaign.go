package svc

import (
	"context"
	"seriouspoop/greedygame/go-common/utils"
	"seriouspoop/greedygame/pkg/model"
	"slices"

	"go.uber.org/zap"
)

func (s *Svc) GetActiveCampaignForDelivery(ctx context.Context, app, os, country string) ([]*model.Campaign, error) {
	ctx, span := s.tracer.Start(ctx, "get_active_campaign_for_delivery")
	defer span.End()

	log := s.logger.Ctx(ctx).With(
		zap.Dict("target",
			zap.String("app", app),
			zap.String("country", country),
			zap.String("os", os)))

	if app == "" || os == "" || country == "" {
		log.Error("empty or invalid app, os or country", zap.Error(ErrImportantFieldMissing))
		return nil, ErrImportantFieldMissing
	}

	targetingRules, err := s.db.GetTargetingRules(ctx)
	if err != nil {
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
		log.Info("no campaign IDs found")
		return nil, ErrNoData
	}

	log.Debug("campaign IDs found for the given target",
		zap.Strings("campaign_ids", utils.StringSlice(campaignIDs)))

	campaigns, err := s.db.GetCampaignFromCIDs(ctx, campaignIDs, model.StatusActive)
	if err != nil {
		return nil, err
	}

	return campaigns, nil
	// c, err := s.db.GetCampaignFromCIDs(ctx, []model.CampaignID{model.CampaignID("668d8555-8021-4448-b2f6-06f7ccfa553e")}, model.StatusActive)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println(c[0].Name)
	// return nil, nil
}

func (s *Svc) matchInInclude(d *model.Dimensions, target string) bool {
	if d == nil {
		return true
	}
	return (len(d.Include) == 0 || slices.Contains(d.Include, target))
}

func (s *Svc) matchInExclude(d *model.Dimensions, target string) bool {
	if d == nil || len(d.Exclude) == 0 {
		return false
	}
	return slices.Contains(d.Exclude, target)
}
