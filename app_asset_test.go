package searchads

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestAppAssetService_List(t *testing.T) {
	client, err := NewV4Client(&http.Client{}, accessToken, &orgID)
	if err != nil {
		panic(err)
	}
	newAG, _, err := client.AppAsset.List(context.Background(), 835599320, &AppAssetOpt{
		CountryOrRegions: []CountryCode{US},
	})
	if err != nil {
		panic(err)
	}
	aa, _ := json.Marshal(newAG)
	fmt.Printf(string(aa))
}
