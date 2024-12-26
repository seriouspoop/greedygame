package handler

import (
	"context"
	"seriouspoop/greedygame/pkg/model"
)

type deliveryServicer interface {
	GetActiveCampaignForDelivery(ctx context.Context, app, os, country string) ([]*model.Campaign, error)
}

type healthServicer interface {
	IsUnhealthy(context.Context) bool
}
