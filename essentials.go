package searchads

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/gbrlsnchs/jwt/v3"
	"net/http"
	"strings"
	"time"
)

type AuthorizationInfo struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
	Scope       string `json:"scope"`
}

func OAuth(clientId, teamId, keyId string, pemDat []byte, expDuration time.Duration) (*AuthorizationInfo, error) {
	block, _ := pem.Decode(pemDat)
	if block == nil {
		return nil, errors.New("pem.Decode error")
	}
	priv, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	alg := jwt.NewES256(jwt.ECDSAPrivateKey(priv))

	IssuedAt := time.Now().UTC()
	ExpirationTime := time.Now().UTC().Add(expDuration)
	pl := jwt.Payload{
		Issuer:         teamId,
		Subject:        clientId,
		Audience:       jwt.Audience{"https://appleid.apple.com"},
		ExpirationTime: jwt.NumericDate(ExpirationTime),
		IssuedAt:       jwt.NumericDate(IssuedAt),
	}
	jwtRe, err := jwt.Sign(pl, alg, jwt.KeyID(keyId))
	cUrl := fmt.Sprintf("https://appleid.apple.com/auth/oauth2/token?grant_type=client_credentials&client_id=%s&client_secret=%s&scope=searchadsorg", clientId, jwtRe)

	payload := strings.NewReader("")
	req, err := http.NewRequest("POST", cUrl, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = CheckResponse(resp)
	if err != nil {
		return nil, err
	}

	authorInfo := AuthorizationInfo{}
	err = json.NewDecoder(resp.Body).Decode(&authorInfo)
	return &authorInfo, nil
}
