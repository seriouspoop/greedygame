package db

import (
	"context"
	"net/url"
	"seriouspoop/greedygame/go-common/db/postgres"
	"seriouspoop/greedygame/pkg/model"
	schema "seriouspoop/greedygame/pkg/repo/db/schema/gen"

	"go.uber.org/zap"
)

var statusModelToDoc = map[model.Status]schema.CampaignStatus{
	model.StatusActive:   schema.CampaignStatusActive,
	model.StatusInactive: schema.CampaignStatusInactive,
}

var statusDocToModel = map[schema.CampaignStatus]model.Status{
	schema.CampaignStatusActive:   model.StatusActive,
	schema.CampaignStatusInactive: model.StatusInactive,
}

func (d *DB) campaignSchemaToModel(ctx context.Context, c schema.Campaign) (*model.Campaign, error) {
	image, err := url.Parse(c.Image)
	if err != nil {
		d.logger.Ctx(ctx).Error("error while image url", zap.Error(err))
		return nil, err
	}
	return &model.Campaign{
		ID:     model.CampaignID(postgres.UUIDToString(c.Cid)),
		Name:   c.Name,
		Image:  image,
		CTA:    c.Cta,
		Status: statusDocToModel[c.Status],
	}, nil
}
