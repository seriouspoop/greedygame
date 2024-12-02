package doc

import (
	"seriouspoop/greedygame/pkg/model"
)

type Dimensions struct {
	Include []string
	Exclude []string
}

type TargetingRuleRec struct {
	CampaignID string
	App        Dimensions
	Country    Dimensions
	OS         Dimensions
}

func fillDimensions(d Dimensions, md **model.Dimensions) {
	if len(d.Exclude) == 0 && len(d.Include) == 0 {
		return
	}
	if (*md) == nil {
		(*md) = &model.Dimensions{}
	}
	if len(d.Exclude) != 0 && len(d.Include) == 0 {
		(*md).Exclude = d.Exclude
	} else if len(d.Exclude) == 0 && len(d.Include) != 0 {
		(*md).Include = d.Include
	} else if len(d.Exclude) != 0 && len(d.Include) != 0 {
		(*md).Exclude = d.Exclude
		(*md).Include = d.Include
	}
}

func (t *TargetingRuleRec) ToModel() *model.TargetingRule {
	targetRuleModel := &model.TargetingRule{CampaignID: model.CampaignID(t.CampaignID)}
	// injecting App
	fillDimensions(t.App, &targetRuleModel.App)
	// injecting Country
	fillDimensions(t.Country, &targetRuleModel.Country)
	// injecting OS
	fillDimensions(t.OS, &targetRuleModel.OS)
	return targetRuleModel
}

var TargetingRuleDummy = []*TargetingRuleRec{
	{
		CampaignID: "spotify",
		App:        Dimensions{},
		Country: Dimensions{
			Include: []string{"us", "canada"},
		},
		OS: Dimensions{},
	},
	{
		CampaignID: "duolingo",
		App:        Dimensions{},
		Country: Dimensions{
			Exclude: []string{"us"},
		},
		OS: Dimensions{
			Include: []string{"android", "ios"},
		},
	},
	{
		CampaignID: "subwaysurfer",
		App: Dimensions{
			Include: []string{"com.gametion.ludokinggame"},
		},
		Country: Dimensions{},
		OS: Dimensions{
			Include: []string{"android"},
		},
	},
}
