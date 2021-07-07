package searchads

import (
	"context"
	"fmt"
)

type AppAssetOpt struct {
	CountryOrRegions []CountryCode `json:"countryOrRegions,omitempty"`
	AssetsGenIds     []string      `json:"assetsGenIds,omitempty"`
}

type CreativeSetAssetsDetail struct {
	CreativeSetDetails map[string]*CreativeSetLocaleDetail `json:"CreativeSetDetails,omitempty"`
}

type CreativeSetLocaleDetail struct {
	LanguageCode               string                                         `json:"languageCode"`
	LanguageDisplayName        string                                         `json:"languageDisplayName"`
	PrimaryLocale              bool                                           `json:"primaryLocale"`
	AppPreviewDeviceWithAssets map[string]*MediaAppPreviewOrScreenshotsDetail `json:"appPreviewDeviceWithAssets,omitempty"`
}

type MediaAppPreviewOrScreenshotsDetail struct {
	DeviceDisplayName           string            `json:"deviceDisplayName"`
	FallBackDevicesDisplayNames map[string]string `json:"fallBackDevicesDisplayNames"`
	Screenshots                 []*Screenshot     `json:"screenshots,omitempty"`
	AppPreviews                 []*AppPreview     `json:"appPreviews,omitempty"`
}

type Screenshot struct {
	AssetGenID   string `json:"assetGenId"`
	AssetURL     string `json:"assetURL"`
	SortPosition int    `json:"sortPosition"`
	SourceHeight int    `json:"sourceHeight"`
	SourceWidth  int    `json:"sourceWidth"`
	Orientation  string `json:"orientation"`
	AssetType    string `json:"assetType"`
}

type AppPreview struct {
	AssetGenId   string `json:"assetGenId"`
	AssetType    string `json:"assetType"`
	AssetURL     string `json:"assetURL"`
	Orientation  string `json:"orientation"`
	SortPosition int64  `json:"sortPosition"`
	SourceHeight int    `json:"sourceHeight"`
	SourceWidth  int    `json:"sourceWidth"`
}

type AppAssetService service

// Fetches assets used with Creative Sets.
func (s *AppAssetService) List(ctx context.Context, adamId int, opt *AppAssetOpt) (*CreativeSetAssetsDetail, *Response, error) {
	req, err := s.client.NewRequest("POST", fmt.Sprintf("creativeappassets/%d", adamId), opt)
	if err != nil {
		return nil, nil, err
	}

	CreativeSet := CreativeSetAssetsDetail{}

	resp, err := s.client.Do(ctx, req, &CreativeSet)
	if err != nil {
		return nil, resp, err
	}

	return &CreativeSet, resp, nil
}

// Fetches supported app preview device size mappings.
func (s *AppAssetService) PreviewDevices(ctx context.Context) (*map[string]string, *Response, error) {
	req, err := s.client.NewRequest("GET", "creativeappmappings/devices", nil)
	if err != nil {
		return nil, nil, err
	}

	DevicesInfo := map[string]string{}

	resp, err := s.client.Do(ctx, req, &DevicesInfo)
	if err != nil {
		return nil, resp, err
	}

	return &DevicesInfo, resp, nil
}
