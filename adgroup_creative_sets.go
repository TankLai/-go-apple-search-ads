package searchads

import (
	"context"
	"fmt"
)

type AdGroupCreativeSetService service

type CreativeSetCreate struct {
	AdamID       int      `json:"adamId,omitempty"`
	Name         string   `json:"name,omitempty"`
	LanguageCode string   `json:"languageCode,omitempty"`
	AssetsGenIds []string `json:"assetsGenIds,omitempty"`
}

type AdGroupCreativeSet struct {
	ID                   int           `json:"id"`
	AdGroupID            int64         `json:"adGroupId"`
	CreativeSetID        int64         `json:"creativeSetId"`
	CampaignID           int64         `json:"campaignId"`
	Status               Status        `json:"status"`
	ServingStatus        ServingStatus `json:"servingStatus"`
	ServingStatusReasons interface{}   `json:"servingStatusReasons"`
	Deleted              bool          `json:"deleted"`
	ModificationTime     string        `json:"modificationTime"`
}

func (s *AdGroupCreativeSetService) Create(ctx context.Context, campaignID, adGroupID int64, data *CreativeSetCreate) (*AdGroupCreativeSet, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}
	if adGroupID == 0 {
		return nil, nil, fmt.Errorf("adGroupID can not be 0")
	}
	reqData := map[string]*CreativeSetCreate{
		"creativeSet": data,
	}
	u := fmt.Sprintf("campaigns/%d/adgroups/%d/adgroupcreativesets/creativesets", campaignID, adGroupID)
	req, err := s.client.NewRequest("POST", u, reqData)
	if err != nil {
		return nil, nil, err
	}

	adGroupCreativeSet := AdGroupCreativeSet{}
	resp, err := s.client.Do(ctx, req, &adGroupCreativeSet)
	if err != nil {
		return nil, resp, err
	}
	return &adGroupCreativeSet, resp, nil
}

// Updates an ad group Creative Set using an identifier.
func (s *AdGroupCreativeSetService) EditStatus(ctx context.Context, campaignID, adGroupID, adGroupCreativeSetID int64, status Status) (*AdGroupCreativeSet, *Response, error) {
	putData := map[string]Status{
		"status": status,
	}
	u := fmt.Sprintf("campaigns/%d/adgroups/%d/adgroupcreativeset/%d", campaignID, adGroupID, adGroupCreativeSetID)
	req, err := s.client.NewRequest("PUT", u, putData)
	if err != nil {
		return nil, nil, err
	}

	adGroupCreativeSet := AdGroupCreativeSet{}
	resp, err := s.client.Do(ctx, req, &adGroupCreativeSet)
	if err != nil {
		return nil, resp, err
	}

	return &adGroupCreativeSet, resp, nil
}

// Creates a Creative Set assignment to an ad group.
func (s *AdGroupCreativeSetService) Assign(ctx context.Context, campaignID, adGroupID, adGroupCreativeSetID int64) (*AdGroupCreativeSet, *Response, error) {
	putData := map[string]int64{
		"creativeSetId": adGroupCreativeSetID,
	}
	u := fmt.Sprintf("campaigns/%d/adgroups/%d/adgroupcreativesets", campaignID, adGroupID)
	req, err := s.client.NewRequest("POST", u, putData)
	if err != nil {
		return nil, nil, err
	}

	c := new(AdGroupCreativeSet)
	resp, err := s.client.Do(ctx, req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, nil
}

// Deletes Creative Sets from a specified ad group.
func (s *AdGroupCreativeSetService) Delete(ctx context.Context, campaignID, adGroupID int64, adGroupCreativeSetIDs []int64) (*int, *Response, error) {
	u := fmt.Sprintf("campaigns/%d/adgroups/%d/adgroupcreativesets/delete/bulk", campaignID, adGroupID)
	req, err := s.client.NewRequest("POST", u, adGroupCreativeSetIDs)
	if err != nil {
		return nil, nil, err
	}

	var c int
	resp, err := s.client.Do(ctx, req, &c)
	if err != nil {
		return nil, resp, err
	}

	return &c, resp, nil
}
