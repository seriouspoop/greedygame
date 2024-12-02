package doc

import (
	"net/url"
	"seriouspoop/greedygame/pkg/model"
)

type CampaignRec struct {
	ID     string
	Name   string
	Image  string
	CTA    string
	Status string
}

func (c *CampaignRec) ToModel() *model.Campaign {
	docToModelStatus := map[string]model.Status{
		"ACTIVE":   model.StatusActive,
		"INACTIVE": model.StatusInactive,
	}

	image, _ := url.Parse(c.Image)
	return &model.Campaign{
		ID:     model.CampaignID(c.ID),
		Name:   c.Name,
		Image:  image,
		CTA:    c.CTA,
		Status: docToModelStatus[c.Status],
	}
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
