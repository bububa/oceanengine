package dpa

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// MetaGetRequest 获取商品库元信息 API Request
type MetaGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PlatformID 商品库id
	PlatformID uint64 `json:"platform_id,omitempty"`
}

// Encode implement GetRequest interface
func (r MetaGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("platform_id", strconv.FormatUint(r.PlatformID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// MetaGetResponse 获取商品库元信息 API Response
type MetaGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// List
		List []Meta `json:"list,omitempty"`
	} `json:"data,omitempty"`
}
