package searchads

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestSearchService_Apps(t *testing.T) {
	client, err := NewV4Client(&http.Client{}, accessToken, &orgID)
	if err != nil {
		panic(err)
	}
	searchApps, _, _ := client.Search.Apps(context.Background(), &SearchAppOptions{
		Query:           "蝉妈妈",
		ReturnOwnedApps: false,
	})
	fmt.Printf("%#v \n", searchApps[0])
}

func TestSearchService_Geo(t *testing.T) {
	client, err := NewV4Client(&http.Client{}, accessToken, &orgID)
	if err != nil {
		panic(err)
	}
	searchApps, _, _ := client.Search.Geo(context.Background(), &SearchGeoOptions{
		Query: "China",
	})
	fmt.Printf("%#v \n", searchApps[0])
}
