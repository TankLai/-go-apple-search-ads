package searchads

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestCampaignNegativeKeywordService_List(t *testing.T) {
	client, err := NewV4Client(&http.Client{}, accessToken, &orgID)
	if err != nil {
		panic(err)
	}
	newK, _, err := client.CampaignNegativeKeyword.List(context.Background(), campaignID, nil)
	if err != nil {
		panic(err)
	}
	aa, _ := json.Marshal(newK)
	fmt.Printf(string(aa))
}

func TestCampaignNegativeKeywordService_CreateBulk(t *testing.T) {
	client, err := NewV4Client(&http.Client{}, accessToken, &orgID)
	if err != nil {
		panic(err)
	}
	newK, _, err := client.CampaignNegativeKeyword.CreateBulk(context.Background(), campaignID, []*NegativeKeyword{
		{
			Text:      "muiscally",
			Status:    KEYWORD_ACTIVE,
			MatchType: MatchTypeExact,
		},
		{
			Text:      "toktok",
			Status:    KEYWORD_ACTIVE,
			MatchType: MatchTypeExact,
		},
	})
	if err != nil {
		panic(err)
	}
	aa, _ := json.Marshal(newK)
	fmt.Printf(string(aa))
}

func TestAdGroupNegativeKeywordService_List(t *testing.T) {
	client, err := NewV4Client(&http.Client{}, accessToken, &orgID)
	if err != nil {
		panic(err)
	}
	newK, _, err := client.AdGroupNegativeKeyword.List(context.Background(), campaignID, 580704484, nil)
	if err != nil {
		panic(err)
	}
	aa, _ := json.Marshal(newK)
	fmt.Printf(string(aa))
}

func TestAdGroupNegativeKeywordService_CreateBulk(t *testing.T) {
	client, err := NewV4Client(&http.Client{}, accessToken, &orgID)
	if err != nil {
		panic(err)
	}
	newK, _, err := client.AdGroupNegativeKeyword.CreateBulk(context.Background(), campaignID, 580704484, []*NegativeKeyword{
		{
			Text:      "muiscally",
			Status:    KEYWORD_ACTIVE,
			MatchType: MatchTypeExact,
		},
		{
			Text:      "toktok",
			Status:    KEYWORD_ACTIVE,
			MatchType: MatchTypeExact,
		},
	})
	if err != nil {
		panic(err)
	}
	aa, _ := json.Marshal(newK)
	fmt.Printf(string(aa))
}

func TestAdGroupNegativeKeywordService_UpdateBulk(t *testing.T) {
	client, err := NewV4Client(&http.Client{}, accessToken, &orgID)
	if err != nil {
		panic(err)
	}
	newK, _, err := client.AdGroupNegativeKeyword.UpdateBulk(context.Background(), campaignID, 580704484, []*NegativeKeyword{
		{
			ID:        580733462,
			Text:      "toktok",
			Status:    KEYWORD_ACTIVE,
			MatchType: MatchTypeExact,
		},
	})
	if err != nil {
		panic(err)
	}
	aa, _ := json.Marshal(newK)
	fmt.Printf(string(aa))
}
