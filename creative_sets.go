package searchads

import (
	"context"
	"fmt"
)

type CreativeSet struct {
	ID                int64               `json:"id,omitempty"`
	OrgID             int64               `json:"orgId,omitempty"`
	AdamID            int64               `json:"adamId,omitempty"`
	Name              string              `json:"name,omitempty"`
	LanguageCode      string              `json:"languageCode,omitempty"`
	Status            Status              `json:"status,omitempty"`
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
func (s *CreativeSetService) List(ctx context.Context, campaignID int64, opt *Selector) ([]*CreativeSet, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}
	mapOpt := map[string]interface{}{}
	if opt == nil {
		mapOpt["selector"] = opt
	}
	req, err := s.client.NewRequest("POST", fmt.Sprintf("campaigns/%d/adgroupcreativesets/find", campaignID), mapOpt)
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
func (s *CreativeSetService) Get(ctx context.Context, creativeSetID int64, opt *AdVariationOpt) (*CreativeSet, *Response, error) {
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
	IncludeDeletedCreativeSetAssets bool      `json:"includeDeletedCreativeSetAssets"`
	Selector                        *Selector `json:"selector"`
}

// Fetches all Creative Sets assigned to an organization.
func (s *CreativeSetService) FindInOrganization(ctx context.Context, opt *OrganizationCreativeSetsOpt) ([]*CreativeSet, *Response, error) {
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

// Updates a Creative Set name using an identifier.
func (s *CreativeSetService) EditName(ctx context.Context, creativeSetId int64, name string) (*CreativeSet, *Response, error) {
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
