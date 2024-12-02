package postgres

import (
	"context"
	"seriouspoop/greedygame/go-common/utils"
	"seriouspoop/greedygame/pkg/model"
	"seriouspoop/greedygame/pkg/svc"
)

var statusModelToString = map[model.Status]string{
	model.StatusUnknown:  "",
	model.StatusActive:   "ACTIVE",
	model.StatusInactive: "INACTIVE",
}

// Get campaigns from cid and status.
// Use StatusUnknown to get all records
func (d *DB) GetCampaignFromCIDs(ctx context.Context, campaignIDs []model.CampaignID, status model.Status) ([]*model.Campaign, error) {
	log := d.logger.WithCtxLogger(ctx)
	if len(campaignIDs) == 0 {
		log.Error().Err(svc.ErrBadInput).Msg("no campaign IDs to process")
		return nil, svc.ErrBadInput
	}

	cidStringSlice := utils.StringSlice(campaignIDs)
	campaigns := []*model.Campaign{}

	if _, ok := statusModelToString[status]; !ok {
		log.Error().Err(svc.ErrBadInput).Msg("invalid status")
		return nil, svc.ErrBadInput
	}

	for _, cid := range cidStringSlice {
		for _, campaignRec := range d.campaignRec {
			if cid == campaignRec.ID && (statusModelToString[status] == campaignRec.Status || status == model.StatusUnknown) {
				campaigns = append(campaigns, campaignRec.ToModel())
			}
		}
	}

	return campaigns, nil
}
