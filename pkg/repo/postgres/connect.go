package postgres

import (
	"seriouspoop/greedygame/go-common/logging"
	"seriouspoop/greedygame/pkg/repo/postgres/doc"
)

type DB struct {
	campaignRec      []*doc.CampaignRec
	targetingRuleRec []*doc.TargetingRuleRec
	logger           *logging.Logger
}

func New(logger *logging.Logger) (*DB, error) {
	return &DB{
		campaignRec:      doc.CampaignDummy,
		targetingRuleRec: doc.TargetingRuleDummy,
		logger:           logger,
	}, nil
}
