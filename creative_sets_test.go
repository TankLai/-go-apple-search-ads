package searchads

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestCreativeSetService_List(t *testing.T) {
	client, err := NewV4Client(&http.Client{}, accessToken, &orgID)
	if err != nil {
		panic(err)
	}
	newCs, _, err := client.CreativeSet.List(context.Background(), campaignID, &Selector{
		Conditions: nil,
		Fields:     nil,
		OrderBy:    nil,
		Pagination: PaginationSelector{
			Offset: 0,
			Limit:  100,
		},
	})
	if err != nil {
		panic(err)
	}
	aa, _ := json.Marshal(newCs)
	fmt.Printf(string(aa))
}

func TestCreativeSetService_FindInOrganization(t *testing.T) {
	client, err := NewV4Client(&http.Client{}, accessToken, &orgID)
	if err != nil {
		panic(err)
	}
	newCs, _, err := client.CreativeSet.FindInOrganization(context.Background(), &OrganizationCreativeSetsOpt{
		Selector: &Selector{
			Conditions: nil,
			Fields:     nil,
			OrderBy:    nil,
			Pagination: PaginationSelector{
				Offset: 0,
				Limit:  100,
			},
		},
		IncludeDeletedCreativeSetAssets: true,
	})
	if err != nil {
		panic(err)
	}
	aa, _ := json.Marshal(newCs)
	fmt.Printf(string(aa))
}

func TestAdGroupCreativeSetService_Create(t *testing.T) {
	client, err := NewV4Client(&http.Client{}, accessToken, &orgID)
	if err != nil {
		panic(err)
	}
	newCs, _, err := client.AdGroupCreativeSet.Create(context.Background(), campaignID, 580704484, &CreativeSetCreate{
		AdamID:       835599320,
		Name:         "test1",
		LanguageCode: "en-US",
		AssetsGenIds: []string{
			"835599320;en-US;9;0;36334a8e0abfbd8ffb43856bf4e54718",
			"835599320;en-US;9;0;dbac55b222a61e1939f19f2640e48dfa",
			"835599320;en-US;9;0;a12241f0e25faa267802493481034799",
			"835599320;en-US;9;0;4a2322d6f9f95df363f5b0aaa1fe666c",
			"835599320;en-US;9;0;8996a6b2625226b09bd25ed7d63c4802",
		},
	})
	if err != nil {
		panic(err)
	}
	aa, _ := json.Marshal(newCs)
	fmt.Printf(string(aa))
}

func TestAdGroupCreativeSetService_Delete(t *testing.T) {
	client, err := NewV4Client(&http.Client{}, accessToken, &orgID)
	if err != nil {
		panic(err)
	}
	newCs, _, err := client.AdGroupCreativeSet.Delete(context.Background(), campaignID, 580704484, []int64{1663745})
	if err != nil {
		panic(err)
	}
	aa, _ := json.Marshal(newCs)
	fmt.Printf(string(aa))
}

func TestAdGroupCreativeSetService_Assign(t *testing.T) {
	client, err := NewV4Client(&http.Client{}, accessToken, &orgID)
	if err != nil {
		panic(err)
	}
	newCs, _, err := client.AdGroupCreativeSet.Assign(context.Background(), campaignID, 580704484, 1663745)
	if err != nil {
		panic(err)
	}
	aa, _ := json.Marshal(newCs)
	fmt.Printf(string(aa))
}
