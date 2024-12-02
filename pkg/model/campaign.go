package model

import "net/url"

type CampaignID string

func (c CampaignID) String() string {
	return string(c)
}

func (c CampaignID) Valid() bool {
	return len(c.String()) > 0
}

type Status int

const (
	StatusUnknown Status = iota
	StatusInactive
	StatusActive
)

type Campaign struct {
	ID     CampaignID
	Name   string
	Image  *url.URL
	CTA    string
	Status Status
}
