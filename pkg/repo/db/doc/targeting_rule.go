package doc

type OS struct {
	Include []string
	Exclude []string
}

type Country struct {
	Include []string
	Exclude []string
}

type App struct {
	Include []string
	Exclude []string
}

type TargetingRuleRec struct {
	CampaignID string
	App        App
	Country    Country
	OS         OS
}

var TargetingRuleDummy = []*TargetingRuleRec{
	{
		CampaignID: "spotify",
		App:        App{},
		Country: Country{
			Include: []string{"US", "Canada"},
		},
		OS: OS{},
	},
	{
		CampaignID: "duolingo",
		App:        App{},
		Country: Country{
			Exclude: []string{"US"},
		},
		OS: OS{
			Include: []string{"Android", "iOS"},
		},
	},
	{
		CampaignID: "subwaysurfer",
		App: App{
			Include: []string{"com.gametion.ludokinggame"},
		},
		Country: Country{},
		OS: OS{
			Include: []string{"Android"},
		},
	},
}
