package handler

import (
	"context"
	"seriouspoop/greedygame/pkg/model"
)

type servicer interface {
	GetCampaignForDelivery(ctx context.Context, app, os, country string) ([]*model.Campaign, error)
}
