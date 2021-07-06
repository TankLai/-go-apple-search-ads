package searchads

import (
	"context"
)

type SearchApp struct {
	AdamID               int64         `json:"adamId"`
	AppName              string        `json:"appName"`
	DeveloperName        string        `json:"developerName"`
	CountryOrRegionCodes []CountryCode `json:"countryOrRegionCodes"`
}

// CampaignService struct to hold individual service
type SearchService service

type SearchAppOptions struct {
	Limit           int    `url:"limit,omitempty"`
	Offset          int    `url:"offset,omitempty"`
	Query           string `url:"query"`
	ReturnOwnedApps bool   `url:"returnOwnedApps,omitempty"`
}

// List function to get Campaigns
func (s *SearchService) Apps(ctx context.Context, opt *SearchAppOptions) ([]*SearchApp, *Response, error) {
	u, err := addOptions("search/apps", opt)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	searchApp := []*SearchApp{}

	resp, err := s.client.Do(ctx, req, &searchApp)
	if err != nil {
		return nil, resp, err
	}

	return searchApp, resp, nil
}

type SearchGeo struct {
	ID          string `json:"id"`
	Entity      string `json:"entity"`
	DisplayName string `json:"displayName"`
}

type SearchGeoOptions struct {
	Limit       int    `url:"limit,omitempty"`
	Offset      int    `url:"offset,omitempty"`
	Query       string `url:"query"`
	CountryCode string `url:"countrycode,omitempty"`
	Entity      string `url:"entity,omitempty"`
}

// List function to get Campaigns
func (s *SearchService) Geo(ctx context.Context, opt *SearchGeoOptions) ([]*SearchGeo, *Response, error) {
	u, err := addOptions("search/geo", opt)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	searchGeo := []*SearchGeo{}

	resp, err := s.client.Do(ctx, req, &searchGeo)
	if err != nil {
		return nil, resp, err
	}

	return searchGeo, resp, nil
}
