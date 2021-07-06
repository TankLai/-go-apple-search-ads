### go-apple-search-ads
Apple Search Ads

Documentation to API V4 can be found [here](https://developer.apple.com/documentation/apple_search_ads "here")


```go
clientId := ""
teamId := ""
keyId := ""
orgID := int64(123456)

pemdat, err := ioutil.ReadFile("../private-key.pem")
if err != nil {
	panic(err)
}

oauthInfo, _ := OAuth(clientId, teamId, keyId, pemdat, 3600 * time.Second)

client := NewV4Client(nil, oauthInfo.AccessToken, &orgID)

fmt.Println(client.AppAsset.PreviewDevices(context.Background()))

```

todo
- test case


Project modified from github.com/jonagold-lab/go-apple-search-ads