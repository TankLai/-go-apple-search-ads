package searchads

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

var (
	clientId    = ""
	teamId      = ""
	keyId       = ""
	orgID       = int64(123123)
	accessToken = ""
)

func TestOAuth(t *testing.T) {
	pemdat, err := ioutil.ReadFile("../private-key.pem")
	if err != nil {
		panic(err)
	}

	oauthInfo, err := OAuth(clientId, teamId, keyId, pemdat, 3600*time.Second)
	if err != nil {
		panic(err)
	}
	fmt.Println(oauthInfo)
	fmt.Println(oauthInfo.AccessToken)
	accessToken := oauthInfo.AccessToken
	client, err := NewV4Client(&http.Client{}, accessToken, &orgID)
	if err != nil {
		panic(err)
	}
	fmt.Println(client.AppAsset.PreviewDevices(context.Background()))
}
