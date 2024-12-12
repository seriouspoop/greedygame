package svc

import (
	"context"
	"seriouspoop/greedygame/go-common/utils"
	"seriouspoop/greedygame/pkg/model"
	"slices"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.uber.org/zap"
)

func (s *Svc) GetActiveCampaignForDelivery(ctx context.Context, app, os, country string) ([]*model.Campaign, error) {
	ctx, span := s.tracer.Start(ctx, "get-active-campaign-for-delivery")
	defer span.End()

	log := s.logger.Ctx(ctx).With(
		zap.Dict("target",
			zap.String("app", app),
			zap.String("country", country),
			zap.String("os", os)))

	if app == "" || os == "" || country == "" {
		log.Error("empty or invalid app, os or country", zap.Error(ErrImportantFieldMissing))
		span.SetStatus(codes.Error, "empty or invalid app, os or country")
		span.RecordError(ErrImportantFieldMissing)
		return nil, ErrImportantFieldMissing
	}

	targetingRules, err := s.db.GetTargetingRules(ctx)
	if err != nil {
		log.Error("error while getting targeting rules", zap.Error(err))
		span.SetStatus(codes.Error, "error while getting targeting rules")
		span.RecordError(err)
		return nil, err
	}

	span.AddEvent("targeting rules retrieved from db")

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
		// do not set span status to error, nothing failed!
		span.RecordError(ErrNoData)
		return nil, ErrNoData
	}

	log.Debug("campaign IDs found for the given target",
		zap.Strings("campaign_ids", utils.StringSlice(campaignIDs)))
	span.SetAttributes(attribute.StringSlice("campaign.ids", utils.StringSlice(campaignIDs)))

	campaigns, err := s.db.GetCampaignFromCIDs(ctx, campaignIDs, model.StatusActive)
	if err != nil {
		log.Error("error while getting campaigns with cids", zap.Error(err))
		span.SetStatus(codes.Error, "error while getting campaigns with cids")
		span.RecordError(err)
		return nil, err
	}

	log.Info("campaigns retrieved from db")
	span.AddEvent("campaigns retrieved from db")
	return campaigns, nil
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
