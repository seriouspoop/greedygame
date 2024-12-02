package model

type Dimensions struct {
	Include []string
	Exclude []string
}

type TargetingRule struct {
	CampaignID CampaignID
	OS         *Dimensions
	Country    *Dimensions
	App        *Dimensions
}
