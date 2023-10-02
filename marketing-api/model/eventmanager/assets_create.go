package eventmanager

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AssetsCreateRequest 创建事件资产 API Request
type AssetsCreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AssetType 资产类型，目前仅支持THIRD_EXTERNAL 三方落地页
	AssetType enum.AssetType `json:"asset_type,omitempty"`
	// ThirdPartAsset 三方落地页资产信息
	ThirdPartAsset *LandingPage `json:"third_part_asset,omitempty"`
	// QuickAppAsset 快应用资产信息
	QuickAppAsset *QuickApp `json:"quick_app_asset,omitempty"`
	// AppAsset 应用信息
	AppAsset *App `json:"app_asset,omitempty"`
	// SiteAsset 橙子落地页信息
	SiteAsset *Site `json:"site_asset,omitempty"`
	// MiniProgramAsset 字节小程序资产信息
	MiniProgramAsset *MiniProgram `json:"mini_program_asset,omitempty"`
}

// Encode implement PostRequest interface
func (r AssetsCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AssetsCreateResponse 创建事件资产 API Response
type AssetsCreateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// AssetID 资产ID
		AssetID uint64 `json:"asset_id,omitempty"`
	} `json:"data,omitempty"`
}
