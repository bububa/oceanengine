package dpa

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AssetsDetailReadRequest 获取投放条件详情 API Request
type AssetsDetailReadRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AssetIDs 物件ID, 最多允许传入100个，可通过【获取投放条件列表】获取
	AssetIDs []uint64 `json:"asset_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r AssetsDetailReadRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	bs, _ := json.Marshal(r.AssetIDs)
	values.Set("asset_ids", string(bs))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AssetsDetailReadResponse 获取投放条件详情 API Response
type AssetsDetailReadResponse struct {
	model.BaseResponse
	Data struct {
		// List 商品库商品列表
		List []Asset `json:"list,omitempty"`
	} `json:"data,omitempty"`
}
