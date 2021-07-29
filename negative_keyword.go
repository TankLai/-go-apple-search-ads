package searchads

import (
	"context"
	"fmt"
)

// NegativeKeyword to define negative Keyword and connection to other
type NegativeKeyword struct {
	ID               int64         `json:"id,omitempty"`
	CampaignID       int64         `json:"campaignId,omitempty"`
	AdGroupID        int64         `json:"adGroupId,omitempty"`
	ModificationTime string        `json:"modificationTime,omitempty"`
	Text             string        `json:"text,omitempty"`
	MatchType        MatchType     `json:"matchType,omitempty"`
	Status           KeywordStatus `json:"status,omitempty"`
	Deleted          bool          `json:"deleted,omitempty"`
}

// CampaignNegativeKeywordService to handle Negative Keywords of
type CampaignNegativeKeywordService service

// List function to get Adgroups from campaign
func (s *CampaignNegativeKeywordService) List(ctx context.Context, campaignID int64, opt *ListOptions) ([]*NegativeKeyword, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}
	u, err := addOptions(fmt.Sprintf("campaigns/%d/negativekeywords", campaignID), opt)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	negativekeywords := []*NegativeKeyword{}
	resp, err := s.client.Do(ctx, req, &negativekeywords)
	if err != nil {
		return nil, resp, err
	}

	return negativekeywords, resp, nil
}

// Fetches negative keywords for campaigns.
func (s *CampaignNegativeKeywordService) Find(ctx context.Context, campaignID int64, selector *Selector) ([]*NegativeKeyword, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("campaigns/%d/negativekeywords/find", campaignID), selector)
	if err != nil {
		return nil, nil, err
	}
	keywords := []*NegativeKeyword{}
	resp, err := s.client.Do(ctx, req, &keywords)
	if err != nil {
		return nil, resp, err
	}

	return keywords, resp, nil
}

// CreateBulk will create multiple Negative Keywords for a campaign
func (s *CampaignNegativeKeywordService) CreateBulk(ctx context.Context, campaignID int64, data []*NegativeKeyword) ([]*NegativeKeyword, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}
	u := fmt.Sprintf("campaigns/%d/negativekeywords/bulk", campaignID)
	req, err := s.client.NewRequest("POST", u, data)
	if err != nil {
		return nil, nil, err
	}
	negativekeywords := []*NegativeKeyword{}
	resp, err := s.client.Do(ctx, req, &negativekeywords)
	if err != nil {
		return nil, resp, err
	}
	return negativekeywords, resp, nil
}

// UpdateBulk will create multiple Negative Keywords for a campaign
func (s *CampaignNegativeKeywordService) UpdateBulk(ctx context.Context, campaignID int64, data []*NegativeKeyword) ([]*NegativeKeyword, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}
	u := fmt.Sprintf("campaigns/%d/negativekeywords/bulk", campaignID)
	req, err := s.client.NewRequest("PUT", u, data)
	if err != nil {
		return nil, nil, err
	}
	negativekeywords := []*NegativeKeyword{}
	resp, err := s.client.Do(ctx, req, &negativekeywords)
	if err != nil {
		return nil, resp, err
	}
	return negativekeywords, resp, nil
}

// Deletes negative keywords from a campaign.
func (s *CampaignNegativeKeywordService) Delete(ctx context.Context, campaignID int64, keywordsIDs []int64) (*int, *Response, error) {
	u := fmt.Sprintf("campaigns/%d/negativekeywords/delete/bulk", campaignID)
	req, err := s.client.NewRequest("POST", u, keywordsIDs)
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

// AdGroupNegativeKeywordService to handle Negative Keywords of
type AdGroupNegativeKeywordService service

// List function to get Adgroups from campaign
func (s *AdGroupNegativeKeywordService) List(ctx context.Context, campaignID int64, adGroupID int64, opt *ListOptions) ([]*NegativeKeyword, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}
	if adGroupID == 0 {
		return nil, nil, fmt.Errorf("adGroupID can not be 0")
	}
	u, err := addOptions(fmt.Sprintf("campaigns/%d/adgroups/%d/negativekeywords", campaignID, adGroupID), opt)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	negativekeywords := []*NegativeKeyword{}
	resp, err := s.client.Do(ctx, req, &negativekeywords)
	if err != nil {
		return nil, resp, err
	}

	return negativekeywords, resp, nil
}

// Fetches negative keywords in ad groups.
func (s *AdGroupNegativeKeywordService) Find(ctx context.Context, campaignID int64, selector *Selector) ([]*NegativeKeyword, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("campaigns/%d/adgroups/negativekeywords/find", campaignID), selector)
	if err != nil {
		return nil, nil, err
	}
	keywords := []*NegativeKeyword{}
	resp, err := s.client.Do(ctx, req, &keywords)
	if err != nil {
		return nil, resp, err
	}

	return keywords, resp, nil
}

// CreateBulk will create multiple Negative Keywords for a campaign
func (s *AdGroupNegativeKeywordService) CreateBulk(ctx context.Context, campaignID, adGroupID int64, data []*NegativeKeyword) ([]*NegativeKeyword, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}
	if adGroupID == 0 {
		return nil, nil, fmt.Errorf("adGroupID can not be 0")
	}
	u := fmt.Sprintf("campaigns/%d/adgroups/%d/negativekeywords/bulk", campaignID, adGroupID)
	req, err := s.client.NewRequest("POST", u, data)
	if err != nil {
		return nil, nil, err
	}
	negativekeywords := []*NegativeKeyword{}
	resp, err := s.client.Do(ctx, req, &negativekeywords)
	if err != nil {
		return nil, resp, err
	}

	return negativekeywords, resp, nil
}

// UpdateBulk will create multiple Negative Keywords for a campaign
func (s *AdGroupNegativeKeywordService) UpdateBulk(ctx context.Context, campaignID, adGroupID int64, data []*NegativeKeyword) ([]*NegativeKeyword, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}
	if adGroupID == 0 {
		return nil, nil, fmt.Errorf("adGroupID can not be 0")
	}
	u := fmt.Sprintf("campaigns/%d/adgroups/%d/negativekeywords/bulk", campaignID, adGroupID)
	req, err := s.client.NewRequest("PUT", u, data)
	if err != nil {
		return nil, nil, err
	}
	negativekeywords := []*NegativeKeyword{}
	resp, err := s.client.Do(ctx, req, &negativekeywords)
	if err != nil {
		return nil, resp, err
	}

	return negativekeywords, resp, nil
}

// Deletes negative keywords from an ad group.
func (s *AdGroupNegativeKeywordService) Delete(ctx context.Context, campaignID, adGroupID int64, keywordsIDs []int64) (*int, *Response, error) {
	u := fmt.Sprintf("campaigns/%d/adgroups/%d/negativekeywords/delete/bulk", campaignID, adGroupID)
	req, err := s.client.NewRequest("POST", u, keywordsIDs)
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
