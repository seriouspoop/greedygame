package postgres

import (
	"context"
	"seriouspoop/greedygame/pkg/model"
	"seriouspoop/greedygame/pkg/svc"
)

func (d *DB) GetTargetingRules(ctx context.Context) ([]*model.TargetingRule, error) {
	log := d.logger.WithCtxLogger(ctx)

	if d.targetingRuleRec == nil || len(d.targetingRuleRec) <= 0 {
		log.Debug().Msg("no items in db or db is nil")
		return nil, svc.ErrNoData
	}

	targetingRules := []*model.TargetingRule{}
	for _, rules := range d.targetingRuleRec {
		targetingRules = append(targetingRules, rules.ToModel())
	}

	log.Debug().Msg("targeting rules retrieved from db")
	return targetingRules, nil
}
