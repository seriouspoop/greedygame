package postgres

import (
	"context"
	"seriouspoop/greedygame/go-common/utils"
	"seriouspoop/greedygame/pkg/model"
	"seriouspoop/greedygame/pkg/svc"

	"go.uber.org/zap"
)

var statusModelToString = map[model.Status]string{
	model.StatusUnknown:  "",
	model.StatusActive:   "ACTIVE",
	model.StatusInactive: "INACTIVE",
}

// Get campaigns from cid and status.
// Use StatusUnknown to get all records
func (d *DB) GetCampaignFromCIDs(ctx context.Context, campaignIDs []model.CampaignID, status model.Status) ([]*model.Campaign, error) {
	log := d.logger.Ctx(ctx)
	if len(campaignIDs) == 0 {
		log.Error("no campaign IDs to process", zap.Error(svc.ErrBadInput))
		return nil, svc.ErrBadInput
	}

	cidStringSlice := utils.StringSlice(campaignIDs)
	campaigns := []*model.Campaign{}

	if _, ok := statusModelToString[status]; !ok {
		log.Error("invalid status", zap.Error(svc.ErrBadInput))
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
