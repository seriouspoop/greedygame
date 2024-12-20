package postgres

import (
	"context"
	"seriouspoop/greedygame/pkg/model"
	"seriouspoop/greedygame/pkg/svc"

	"go.uber.org/zap"
)

func (d *DB) GetTargetingRules(ctx context.Context) ([]*model.TargetingRule, error) {
	log := d.logger.Ctx(ctx)

	if d.targetingRuleRec == nil || len(d.targetingRuleRec) <= 0 {
		log.Error("no items in db or db is nil", zap.Error(svc.ErrNoData))
		return nil, svc.ErrNoData
	}

	targetingRules := []*model.TargetingRule{}
	for _, rules := range d.targetingRuleRec {
		targetingRules = append(targetingRules, rules.ToModel())
	}

	log.Debug("targeting rules retrieved from db", zap.Any("targeting rules", targetingRules))
	return targetingRules, nil
}
