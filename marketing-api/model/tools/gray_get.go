package tools

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GrayGetRequest 查询白名单能力 API Request
type GrayGetRequest struct {
	// AdvertiserID 广告账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// GrayKeys 白名单能力key，目前仅支持单次查询1个白名单能力key
	GrayKeys []enum.GrayKey `json:"gray_keys,omitempty"`
}

// Encode implement GetRequest interface
func (r GrayGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("gray_keys", string(util.JSONMarshal(r.GrayKeys)))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// GrayGetResponse 查询白名单能力 API Response
type GrayGetResponse struct {
	model.BaseResponse
	Data struct {
		// List 命中白名单请求
		List []GrayItem `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

// GrayItem 白名单
type GrayItem struct {
	// GrayKey 白名单能力的唯一key
	GrayKey string `json:"gray_key,omitempty"`
	// InWhitelist 是否命中白名单。0代表命中，1代表未命中
	InWhitelist int `json:"in_whitelist,omitempty"`
}
