package db

import (
	"seriouspoop/greedygame/go-common/db/postgres"
	"seriouspoop/greedygame/pkg/model"
	schema "seriouspoop/greedygame/pkg/repo/db/schema/gen"
)

func (d *DB) toDimension(inc []string, ex []string) *model.Dimensions {
	return &model.Dimensions{
		Include: inc,
		Exclude: ex,
	}
}

func (d *DB) targetingSchemaToModel(t schema.TargetingRule) *model.TargetingRule {
	return &model.TargetingRule{
		CampaignID: model.CampaignID(postgres.UUIDToString(t.Cid)),
		OS:         d.toDimension(t.OsInclude, t.OsExclude),
		App:        d.toDimension(t.AppInclude, t.AppExclude),
		Country:    d.toDimension(t.CountryInclude, t.CountryExclude),
	}
}
