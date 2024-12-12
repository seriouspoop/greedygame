package svc

import (
	"context"
	"net/url"
	"seriouspoop/greedygame/go-common/logging"
	"seriouspoop/greedygame/go-common/observer"
	"seriouspoop/greedygame/pkg/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func buildUrl(urlString string) *url.URL {
	url, err := url.Parse(urlString)
	if err != nil {
		return nil
	}
	return url
}

func TestGetActiveCampaignForDelivery(t *testing.T) {
	tests := []struct {
		name          string
		app           string
		os            string
		country       string
		targetingRule []*model.TargetingRule
		campaignIDs   []model.CampaignID
		campaigns     []*model.Campaign
		setMocks      func(ctx context.Context, db *MockdbHelper, targetingRules []*model.TargetingRule, campaignIDs []model.CampaignID, campaigns []*model.Campaign, err error)
		err           error
	}{
		{
			name:          "invalid query paramter, expect error",
			app:           "",
			os:            "android",
			country:       "india",
			targetingRule: nil,
			campaignIDs:   nil,
			campaigns:     nil,
			setMocks: func(ctx context.Context, db *MockdbHelper, targetingRules []*model.TargetingRule, campaignIDs []model.CampaignID, campaigns []*model.Campaign, err error) {
			},
			err: ErrImportantFieldMissing,
		},
		{
			name:    "valid query parameters, but no target match, expect error",
			app:     "random.app.com",
			os:      "ios",
			country: "india",
			targetingRule: []*model.TargetingRule{
				{
					CampaignID: "spotify",
					App:        nil,
					Country: &model.Dimensions{
						Include: []string{"us", "canada"},
					},
					OS: nil,
				},
				{
					CampaignID: "duolingo",
					App:        nil,
					Country: &model.Dimensions{
						Exclude: []string{"us"},
					},
					OS: &model.Dimensions{
						// ios missing here, hence no match found expected
						Include: []string{"android"},
					},
				},
				{
					CampaignID: "subwaysurfer",
					App: &model.Dimensions{
						Include: []string{"com.gametion.ludokinggame"},
					},
					Country: nil,
					OS: &model.Dimensions{
						Include: []string{"android"},
					},
				},
			},
			campaignIDs: []model.CampaignID{},
			campaigns:   nil,
			setMocks: func(ctx context.Context, db *MockdbHelper, targetingRules []*model.TargetingRule, campaignIDs []model.CampaignID, campaigns []*model.Campaign, err error) {
				db.EXPECT().GetTargetingRules(ctx).Times(1).Return(targetingRules, nil)
			},
			err: ErrNoData,
		},
		{
			name:    "valid query parameters, target match, expect no error",
			app:     "com.gametion.ludokinggame",
			os:      "android",
			country: "us",
			targetingRule: []*model.TargetingRule{
				{
					CampaignID: "spotify",
					App:        nil,
					Country: &model.Dimensions{
						Include: []string{"us", "canada"},
					},
					OS: nil,
				},
				{
					CampaignID: "duolingo",
					App:        nil,
					Country: &model.Dimensions{
						Exclude: []string{"us"},
					},
					OS: &model.Dimensions{
						// ios present here, hence match found expected
						Include: []string{"android", "ios"},
					},
				},
				{
					CampaignID: "subwaysurfer",
					App: &model.Dimensions{
						Include: []string{"com.gametion.ludokinggame"},
					},
					Country: nil,
					OS: &model.Dimensions{
						Include: []string{"android"},
					},
				},
			},
			campaignIDs: []model.CampaignID{model.CampaignID("spotify"), model.CampaignID("subwaysurfer")},
			campaigns: []*model.Campaign{
				{
					ID:     model.CampaignID("spotify"),
					Name:   "Spotify - Music for everyone",
					Image:  buildUrl("https://somelink"),
					CTA:    "Download",
					Status: model.StatusActive,
				},
				{
					ID:     "subwaysurfer",
					Name:   "Subway Surfer",
					Image:  buildUrl("https://somelink3"),
					CTA:    "Play",
					Status: model.StatusActive,
				},
			},
			setMocks: func(ctx context.Context, db *MockdbHelper, targetingRules []*model.TargetingRule, campaignIDs []model.CampaignID, campaigns []*model.Campaign, err error) {
				db.EXPECT().GetTargetingRules(ctx).Times(1).Return(targetingRules, nil)
				db.EXPECT().GetCampaignFromCIDs(ctx, campaignIDs, model.StatusActive).Times(1).Return(campaigns, nil)
			},
			err: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.Background()
			tracer := observer.NewNoopTracer()
			mockCtx := tracer.MockContext(ctx)
			logger := logging.NewTestLogger()
			db := NewMockdbHelper(ctrl)

			s := New(db, logger, tracer)
			test.setMocks(mockCtx, db, test.targetingRule, test.campaignIDs, test.campaigns, test.err)
			campaigns, err := s.GetActiveCampaignForDelivery(ctx, test.app, test.os, test.country)

			assert.Equal(t, test.err, err)
			assert.Equal(t, test.campaigns, campaigns)
		})
	}
}
