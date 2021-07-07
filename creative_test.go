package searchads

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestCreativeSetService_Create(t *testing.T) {
	client, err := NewV4Client(&http.Client{}, accessToken, &orgID)
	if err != nil {
		panic(err)
	}
	newCs, _, err := client.CreativeSet.Create(context.Background(), campaignID, 580704484, &CreativeSetCreate{
		AdamID:       835599320,
		Name:         "test1",
		LanguageCode: "en-US",
		AssetsGenIds: []string{
			"835599320;en-US;9;0;36334a8e0abfbd8ffb43856bf4e54718",
			"835599320;en-US;9;0;dbac55b222a61e1939f19f2640e48dfa",
			"835599320;en-US;9;0;a12241f0e25faa267802493481034799",
		},
	})
	if err != nil {
		panic(err)
	}
	aa, _ := json.Marshal(newCs)
	fmt.Printf(string(aa))
}

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
