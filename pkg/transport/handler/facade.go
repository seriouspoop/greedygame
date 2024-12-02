package handler

import (
	"context"
	"seriouspoop/greedygame/pkg/model"
)

type servicer interface {
	GetActiveCampaignForDelivery(ctx context.Context, app, os, country string) ([]*model.Campaign, error)
}
