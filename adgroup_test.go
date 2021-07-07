package searchads

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"
)

var campaignID = int64(580657822)

func TestAdGroupService_List(t *testing.T) {
	client, err := NewV4Client(&http.Client{}, accessToken, &orgID)
	if err != nil {
		panic(err)
	}
	newAG, _, err := client.AdGroup.List(context.Background(), campaignID, &ListOptions{
		Limit:  100,
		Offset: 0,
	})
	if err != nil {
		panic(err)
	}
	aa, _ := json.Marshal(newAG)
	fmt.Printf(string(aa))
}

func TestAdGroupService_Get(t *testing.T) {
	client, err := NewV4Client(&http.Client{}, accessToken, &orgID)
	if err != nil {
		panic(err)
	}
	newAG, _, err := client.AdGroup.Get(context.Background(), campaignID, 580704484)
	if err != nil {
		panic(err)
	}
	aa, _ := json.Marshal(newAG)
	fmt.Printf(string(aa))
}

func TestAdGroupService_Create(t *testing.T) {
	client, err := NewV4Client(&http.Client{}, accessToken, &orgID)
	if err != nil {
		panic(err)
	}
	newAG, _, err := client.AdGroup.Create(context.Background(), campaignID, &AdGroup{
		CampaignID: campaignID,
		Name:       "go-ad-group-example",
		// CpaGoal:                nil,
		AutomatedKeywordsOptIn: false,
		DefaultBidAmount: &Amount{
			Amount:   "0.5",
			Currency: "USD",
		},
		StartTime:    time.Now().Add(5 * time.Hour).Format("2006-01-02T15:04:05.000"),
		EndTime:      time.Now().Add(15 * time.Hour).Format("2006-01-02T15:04:05.000"),
		OrgID:        orgID,
		PricingModel: CPC,
		TargetingDimensions: &TargetingDimensions{
			Age: Age{
				Included: []AgeObj{
					{
						MaxAge: 40,
						MinAge: 20,
					},
				},
			},
		},
		Status: PAUSED,
	})
	if err != nil {
		panic(err)
	}
	aa, _ := json.Marshal(newAG)
	fmt.Printf(string(aa))
}

func TestAdGroupService_Edit(t *testing.T) {
	client, err := NewV4Client(&http.Client{}, accessToken, &orgID)
	if err != nil {
		panic(err)
	}
	newAG, _, err := client.AdGroup.Edit(context.Background(), campaignID, 580704484, &AdGroup{
		Name: "go-ad-group-example-up",
		// CpaGoal:                nil,
		AutomatedKeywordsOptIn: false,
		DefaultBidAmount: &Amount{
			Amount:   "0.5",
			Currency: "USD",
		},
		StartTime: time.Now().Add(5 * time.Hour).Format("2006-01-02T15:04:05.000"),
		EndTime:   time.Now().Add(15 * time.Hour).Format("2006-01-02T15:04:05.000"),
		TargetingDimensions: &TargetingDimensions{
			Age: Age{
				Included: []AgeObj{
					{
						MaxAge: 45,
						MinAge: 20,
					},
				},
			},
		},
		Status: PAUSED,
	})
	if err != nil {
		panic(err)
	}
	aa, _ := json.Marshal(newAG)
	fmt.Printf(string(aa))
}

func TestAdGroupService_Delete(t *testing.T) {
	// 580704484
	client, err := NewV4Client(&http.Client{}, accessToken, &orgID)
	if err != nil {
		panic(err)
	}
	resp, err := client.AdGroup.Delete(context.Background(), campaignID, 580683578)
	if err != nil {
		panic(err)
	}
	aa, _ := json.Marshal(resp)
	fmt.Printf(string(aa))
}
