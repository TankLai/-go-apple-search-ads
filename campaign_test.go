package searchads

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestCampaignService_List(t *testing.T) {
	client, err := NewV4Client(&http.Client{}, accessToken, &orgID)
	if err != nil {
		panic(err)
	}
	newC, _, err := client.Campaign.List(context.Background(), &ListOptions{
		Limit:  100,
		Offset: 0,
	})
	if err != nil {
		panic(err)
	}
	aa, _ := json.Marshal(newC)
	fmt.Printf(string(aa))
}

func TestCampaignService_Get(t *testing.T) {
	client, err := NewV4Client(&http.Client{}, accessToken, &orgID)
	if err != nil {
		panic(err)
	}
	newC, _, err := client.Campaign.Get(context.Background(), 580657822)
	if err != nil {
		panic(err)
	}
	// ID:580657822
	aa, _ := json.Marshal(newC)
	fmt.Printf(string(aa))
}

func TestCampaignService_Create(t *testing.T) {
	client, err := NewV4Client(&http.Client{}, accessToken, &orgID)
	if err != nil {
		panic(err)
	}
	newC, _, err := client.Campaign.Create(context.Background(), &Campaign{
		OrgID: orgID,
		Name:  "go-sdk-example2",
		BudgetAmount: &Amount{
			Amount:   "10",
			Currency: "USD",
		},
		DailyBudgetAmount: &Amount{
			Amount:   "1",
			Currency: "USD",
		},
		AdamID:             1508944184,
		CountriesOrRegions: []CountryCode{US},
		AdChannelType:      SEARCH,
		BillingEvent:       TAPS,
		SupplySources:      []SupplySources{APPSTORE_SEARCH_RESULTS},
	})
	if err != nil {
		panic(err)
	}
	// ID:580657822
	aa, _ := json.Marshal(newC)
	fmt.Printf(string(aa))
}

func TestCampaignService_Edit(t *testing.T) {
	client, err := NewV4Client(&http.Client{}, accessToken, &orgID)
	if err != nil {
		panic(err)
	}
	newC, _, err := client.Campaign.Edit(context.Background(), 580657822, &Campaign{
		Name: "go-sdk-example-update",
		BudgetAmount: &Amount{
			Amount:   "15",
			Currency: "USD",
		},
		DailyBudgetAmount: &Amount{
			Amount:   "1",
			Currency: "USD",
		},
		CountriesOrRegions: []CountryCode{US, FR},
	}, true)
	if err != nil {
		panic(err)
	}
	// ID:580657822
	aa, _ := json.Marshal(newC)
	fmt.Printf(string(aa))
}

func TestCampaignService_Delete(t *testing.T) {
	// 580564833
	client, err := NewV4Client(&http.Client{}, accessToken, &orgID)
	if err != nil {
		panic(err)
	}
	resp, err := client.Campaign.Delete(context.Background(), 580564833)
	if err != nil {
		panic(err)
	}
	aa, _ := json.Marshal(resp)
	fmt.Printf(string(aa))
}

func TestCountryCode_String(t *testing.T) {
	source := []CountryCode{US, CN}
	bytes, _ := json.Marshal(source)
	fmt.Println(string(bytes))
	sStr := make([]string, 0)
	json.Unmarshal(bytes, &sStr)
	fmt.Println(sStr)
	ss := strings.Join(sStr, ",")
	fmt.Println(ss)
}
