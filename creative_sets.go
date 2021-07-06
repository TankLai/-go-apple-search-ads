package searchads

import (
	"context"
	"fmt"
)

type CreativeSet struct {
	ID                int                 `json:"id,omitempty"`
	OrgID             int                 `json:"orgId,omitempty"`
	AdamID            int                 `json:"adamId,omitempty"`
	Name              string              `json:"name,omitempty"`
	LanguageCode      string              `json:"languageCode,omitempty"`
	Status            string              `json:"status,omitempty"`
	StatusReasons     []string            `json:"statusReasons,omitempty"`
	CreativeSetAssets []*CreativeSetAsset `json:"creativeSetAssets,omitempty"`
}

type CreativeSetAsset struct {
	ID    int   `json:"id"`
	Asset Asset `json:"asset"`
}

type Asset struct {
	AssetGenID       string `json:"assetGenId"`
	Type             string `json:"type"`
	AppPreviewDevice string `json:"appPreviewDevice"`
	Orientation      string `json:"orientation"`
	Deleted          bool   `json:"deleted"`
}

type CreativeSetService service

// Fetches all Creative Sets assigned to ad groups.
func (s *CreativeSetService) List(ctx context.Context, campaignID int, opt *Selector) ([]*CreativeSet, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}
	req, err := s.client.NewRequest("POST", fmt.Sprintf("campaigns/%d/adgroupcreativesets/find", campaignID), opt)
	if err != nil {
		return nil, nil, err
	}

	CreativeSets := []*CreativeSet{}

	resp, err := s.client.Do(ctx, req, &CreativeSets)
	if err != nil {
		return nil, resp, err
	}

	return CreativeSets, resp, nil
}

type AdVariationOpt struct {
	IncludeDeletedCreativeSetAssets bool `json:"includeDeletedCreativeSetAssets"`
}

// Fetches supported app preview device size mappings.
func (s *CreativeSetService) Get(ctx context.Context, creativeSetID int, opt *AdVariationOpt) (*CreativeSet, *Response, error) {
	if creativeSetID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}
	u, err := addOptions(fmt.Sprintf("creativesets/%d", creativeSetID), opt)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	creativeSet := CreativeSet{}

	resp, err := s.client.Do(ctx, req, &creativeSet)
	if err != nil {
		return nil, resp, err
	}

	return &creativeSet, resp, nil
}

type OrganizationCreativeSetsOpt struct {
	IncludeDeletedCreativeSetAssets bool     `json:"includeDeletedCreativeSetAssets"`
	Selector                        Selector `json:"selector"`
}

// Fetches all Creative Sets assigned to an organization.
func (s *CreativeSetService) Find(ctx context.Context, opt *Selector) ([]*CreativeSet, *Response, error) {
	req, err := s.client.NewRequest("POST", "creativesets/find", opt)
	if err != nil {
		return nil, nil, err
	}

	CreativeSets := []*CreativeSet{}

	resp, err := s.client.Do(ctx, req, &CreativeSets)
	if err != nil {
		return nil, resp, err
	}

	return CreativeSets, resp, nil
}

type CreativeSetCreate struct {
	AdamID       int      `json:"adamId,omitempty"`
	Name         string   `json:"name,omitempty"`
	LanguageCode string   `json:"languageCode,omitempty"`
	AssetsGenIds []string `json:"assetsGenIds,omitempty"`
}

type AdGroupCreativeSet struct {
	ID                   int            `json:"id"`
	AdGroupID            int            `json:"adGroupId"`
	CreativeSetID        int            `json:"creativeSetId"`
	CampaignID           int            `json:"campaignId"`
	Status               *Status        `json:"status"`
	ServingStatus        *ServingStatus `json:"servingStatus"`
	ServingStatusReasons []string       `json:"servingStatusReasons"`
	Deleted              bool           `json:"deleted"`
	ModificationTime     string         `json:"modificationTime"`
}

func (s *CreativeSetService) Create(ctx context.Context, campaignID, adGroupID int, data *CreativeSetCreate) (*AdGroupCreativeSet, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}
	if adGroupID == 0 {
		return nil, nil, fmt.Errorf("adGroupID can not be 0")
	}
	u := fmt.Sprintf("campaigns/%d/adgroups/%d/adgroupcreativesets/creativesets", campaignID, adGroupID)
	req, err := s.client.NewRequest("POST", u, data)
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
func (s *CreativeSetService) EditStatus(ctx context.Context, campaignID, adGroupID, adGroupCreativeSetID int, status Status) (*AdGroupCreativeSet, *Response, error) {
	putData := map[string]Status{
		"status": status,
	}
	u := fmt.Sprintf("campaigns/%d/adgroups/%d/adgroupcreativeset/%d", campaignID, adGroupID, adGroupCreativeSetID)
	req, err := s.client.NewRequest("PUT", u, putData)
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
func (s *CreativeSetService) Delete(ctx context.Context, campaignID, adGroupID int, adGroupCreativeSetIDs []int) (*int, *Response, error) {
	u := fmt.Sprintf("campaigns/%d/adgroups/%d/adgroupcreativesets/delete/bulk", campaignID, adGroupID)
	req, err := s.client.NewRequest("PUT", u, adGroupCreativeSetIDs)
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

// Updates a Creative Set name using an identifier.
func (s *CreativeSetService) EditName(ctx context.Context, creativeSetId int, name string) (*CreativeSet, *Response, error) {
	putData := map[string]string{
		"name": name,
	}
	u := fmt.Sprintf("creativesets/%d", creativeSetId)
	req, err := s.client.NewRequest("PUT", u, putData)
	if err != nil {
		return nil, nil, err
	}

	c := new(CreativeSet)
	resp, err := s.client.Do(ctx, req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, nil
}
