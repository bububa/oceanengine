package event

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AssetsCreateRequest 创建资产 API Request
type AssetsCreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AssetType 资产类型，目前仅支持THIRD_EXTERNAL 三方落地页
	AssetType enum.AssetType `json:"asset_type,omitempty"`
	// ThirdPartAsset 三方落地页资产信息
	ThirdPartAsset *ThirdPartAsset `json:"third_part_asset,omitempty"`
}

// ThirdPartAsset 三方落地页资产信息
type ThirdPartAsset struct {
	// Name 落地页名称，长度限制为25，一个字符长度为1
	Name string `json:"name,omitempty"`
	// Description 落地页名称，长度限制为150，一个字符长度为1
	Description string `json:"description,omitempty"`
}

// Encode implement PostRequest interface
func (r AssetsCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AssetsCreateResponse 创建资产 API Response
type AssetsCreateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// AssetID 资产ID
		AssetID uint64 `json:"asset_id,omitempty"`
	} `json:"data,omitempty"`
}
