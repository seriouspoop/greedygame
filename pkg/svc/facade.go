package svc

import (
	"context"
	"seriouspoop/greedygame/pkg/model"
)

type dbHelper interface {
	GetTargetingRules(ctx context.Context) ([]*model.TargetingRule, error)
}
