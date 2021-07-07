package searchads

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestAdGroupTargetingKeywordService_List(t *testing.T) {
	client, err := NewV4Client(&http.Client{}, accessToken, &orgID)
	if err != nil {
		panic(err)
	}
	newK, _, err := client.AdGroupTargetingKeyword.List(context.Background(), campaignID, 580704484, nil)
	if err != nil {
		panic(err)
	}
	aa, _ := json.Marshal(newK)
	fmt.Printf(string(aa))
}
func TestAdGroupTargetingKeywordService_CreateBulk(t *testing.T) {
	client, err := NewV4Client(&http.Client{}, accessToken, &orgID)
	if err != nil {
		panic(err)
	}
	newK, _, err := client.AdGroupTargetingKeyword.CreateBulk(context.Background(), campaignID, 580704484, []*TargetingKeyword{
		{
			Text:      "muiscally",
			Status:    KEYWORD_ACTIVE,
			MatchType: MatchTypeExact,
			BidAmount: Amount{
				Amount:   "0.5",
				Currency: "USD",
			},
		},
		{
			Text:      "toktok",
			Status:    KEYWORD_ACTIVE,
			MatchType: MatchTypeExact,
			BidAmount: Amount{
				Amount:   "0.6",
				Currency: "USD",
			},
		},
	})
	if err != nil {
		panic(err)
	}
	aa, _ := json.Marshal(newK)
	fmt.Printf(string(aa))
}
