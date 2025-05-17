package eventmanager

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AllAssetsDetailRequest 获取已创建资产详情（新） API Request
type AllAssetsDetailRequest struct {
	// AssetIDs 资产id列表，list长度最长50
	// 当账户下不存在该资产id时不会返回详情信息；当资产共享失效时，不会返回详情信息。
	AssetIDs []uint64 `json:"asset_ids,omitempty"`
	// AdvertiserID 广告主 id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implements GetRequest interface
func (r AllAssetsDetailRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("asset_ids", string(util.JSONMarshal(r.AssetIDs)))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AllAssetsDetailResponse 获取已创建资产详情（新） API Response
type AllAssetsDetailResponse struct {
	model.BaseResponse
	Data struct {
		// AssetList 资产列表
		AssetList []AssetDetail `json:"asset_list,omitempty"`
	} `json:"data,omitempty"`
}
