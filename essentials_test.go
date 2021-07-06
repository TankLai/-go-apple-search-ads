package searchads

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

func TestOAuth(t *testing.T) {
	clientId := ""
	teamId := ""
	keyId := ""

	pemdat, err := ioutil.ReadFile("../private-key.pem")
	if err != nil {
		panic(err)
	}

	fmt.Println(OAuth(clientId, teamId, keyId, pemdat, 3600*time.Second))
}
