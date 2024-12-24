package db

import (
	"context"
	"seriouspoop/greedygame/go-common/db/postgres"
	"seriouspoop/greedygame/pkg/model"
	"seriouspoop/greedygame/pkg/svc"

	"go.uber.org/zap"
)

// Get campaigns from cid and status.
// Use StatusUnknown to get all records
func (d *DB) GetCampaignFromCIDs(ctx context.Context, campaignIDs []model.CampaignID, status model.Status) ([]*model.Campaign, error) {
	ctx, span := d.tracer.Start(ctx, "[DB]_get_campaign_from_cids_db")
	defer span.End()

	log := d.logger.Ctx(ctx).With(zap.Stringers("campaign_ids", campaignIDs))

	if len(campaignIDs) == 0 {
		log.Error("no campaign IDs to process", zap.Error(svc.ErrBadInput))
		return nil, svc.ErrBadInput
	}

	if _, ok := statusModelToDoc[status]; !ok {
		log.Error("invalid status", zap.Error(svc.ErrBadInput))
		return nil, svc.ErrBadInput
	}

	pgUUIDs, err := postgres.ToUUIDs(campaignIDs)
	if err != nil {
		log.Error("error while converting CIDs to UUIDs", zap.Error(err))
		return nil, svc.ErrBadInput
	}

	campaignsRec, err := d.query.GetCampaignFromCIDs(ctx, pgUUIDs)
	if err != nil {
		log.Error("error while getting campaigns from db", zap.Error(err))
		// Always return db errors as svc errors
		return nil, postgres.SvcError(err)
	}

	if len(campaignsRec) == 0 {
		log.Debug("no campaigns found with cids")
		return nil, svc.ErrNoData
	}

	campaigns := []*model.Campaign{}
	for _, campaign := range campaignsRec {
		campaignModel, err := d.campaignSchemaToModel(ctx, campaign)
		if err != nil {
			return nil, svc.ErrBadInput
		}
		campaigns = append(campaigns, campaignModel)
	}

	log.Info("fetched campaigns from db")
	return campaigns, nil
}
