package dpa

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CategoryGetRequest 获取DPA分类 API Request
type CategoryGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PlatformID 商品库id
	PlatformID uint64 `json:"platform_id,omitempty"`
}

// Encode implement GetRequest interface
func (r CategoryGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("platform_id", strconv.FormatUint(r.PlatformID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// CategoryGetResponse 获取DPA分类 API Response
type CategoryGetResponse struct {
	model.BaseResponse
	Data struct {
		// List DPA分类列表
		List []Category `json:"list,omitempty"`
	} `json:"data,omitempty"`
}
