package db

import (
	"seriouspoop/greedygame/pkg/repo/db/doc"

	"github.com/rs/zerolog"
)

type DB struct {
	campaignRec      []*doc.CampaignRec
	targetingRuleRec []*doc.TargetingRuleRec
	logger           *zerolog.Logger
}

func New(logger *zerolog.Logger) (*DB, error) {
	return &DB{
		campaignRec:      doc.CampaignDummy,
		targetingRuleRec: doc.TargetingRuleDummy,
		logger:           logger,
	}, nil
}
