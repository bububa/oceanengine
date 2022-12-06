package dpa

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DictGetRequest 获取DPA词包 API Request
type DictGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PlatformID 商品库id
	PlatformID uint64 `json:"platform_id,omitempty"`
}

// Encode implement GetRequest interface
func (r DictGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("platform_id", strconv.FormatUint(r.PlatformID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// DictGetResponse 获取DPA词包 API Response
type DictGetResponse struct {
	model.BaseResponse
	Data struct {
		// List DPA词包列表
		List []Dict `json:"list,omitempty"`
	} `json:"data,omitempty"`
}
