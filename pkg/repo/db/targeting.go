package db

import (
	"context"
	"seriouspoop/greedygame/go-common/db/postgres"
	"seriouspoop/greedygame/pkg/model"
	"seriouspoop/greedygame/pkg/svc"

	"go.uber.org/zap"
)

func (d *DB) GetTargetingRules(ctx context.Context) ([]*model.TargetingRule, error) {
	ctx, span := d.tracer.Start(ctx, "[DB]_get_targeting_rules_from")
	defer span.End()

	log := d.logger.Ctx(ctx)

	rules, err := d.query.GetAllTargetingRules(ctx)
	if err != nil {
		log.Error("error while getting targeting rules from db", zap.Error(err))
		return nil, postgres.SvcError(err)
	}

	// sqlc doesn't return ErrNoRows for queries with :many
	if len(rules) == 0 {
		log.Info("no targeting rules present in db")
		return nil, svc.ErrNoData
	}

	targetingRules := []*model.TargetingRule{}
	for _, rule := range rules {
		targetingRules = append(targetingRules, d.targetingSchemaToModel(rule))
	}

	log.Info("targeting rules retrieved from db")
	return targetingRules, nil
}
