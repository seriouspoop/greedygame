package doc

type CampaignRec struct {
	ID     string
	Name   string
	Image  string
	CTA    string
	Status string
}

var CampaignDummy = []*CampaignRec{
	{
		ID:     "spotify",
		Name:   "Spotify - Music for everyone",
		Image:  "https://somelink",
		CTA:    "Download",
		Status: "ACTIVE",
	},
	{
		ID:     "duolingo",
		Name:   "Duolingo: Best way to learn",
		Image:  "https://somelink2",
		CTA:    "Install",
		Status: "ACTIVE",
	},
	{
		ID:     "subwaysurfer",
		Name:   "Subway Surfer",
		Image:  "https://somelink3",
		CTA:    "Play",
		Status: "ACTIVE",
	},
}
