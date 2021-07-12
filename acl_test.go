package searchads

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestACLService_List(t *testing.T) {
	client, err := NewV4Client(&http.Client{}, accessToken, &orgID)
	if err != nil {
		panic(err)
	}
	newAG, _, err := client.ACL.List(context.Background(), &ListOptions{
		Limit:  100,
		Offset: 0,
	})
	if err != nil {
		panic(err)
	}
	aa, _ := json.Marshal(newAG)
	fmt.Printf(string(aa))
}

func TestACLService_Me(t *testing.T) {
	client, err := NewV4Client(&http.Client{}, accessToken, &orgID)
	if err != nil {
		panic(err)
	}
	newAG, _, err := client.ACL.Me(context.Background())
	if err != nil {
		panic(err)
	}
	aa, _ := json.Marshal(newAG)
	fmt.Printf(string(aa))
}
