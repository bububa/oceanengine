package advertiser

import (
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// TypeGetRequest 获取千川账户类型 API Request
type TypeGetRequest struct {
	// AdvertiserIDs 千川广告主账户id，一次请求不超过20个
	AdvertiserIDs []uint64 `json:"advertiser_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r TypeGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_ids", string(util.JSONMarshal(r.AdvertiserIDs)))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// TypeGetResponse 获取千川账户类型 API Response
type TypeGetResponse struct {
	model.BaseResponse
	Data struct {
		// List 广告主数据列表
		List []Advertiser `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

type Advertiser struct {
	// EcpType 账户类型，可选值:
	// SHOP: 商家
	// SHOP_STAR: 商家达人
	// COMMON_STAR: 普通达人
	// AGENT: 百应机构
	EcpType qianchuan.EcpType `json:"ecp_type,omitempty"`
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}
