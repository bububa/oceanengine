package tools

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ActionTextGetRequest 行动号召字段内容获取 API Request
type ActionTextGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// LandingType 推广类型
	// 允许值：APP，SHOP，LINK
	LandingType enum.LandingType `json:"landing_type,omitempty"`
	// AdvanceCreativeType 附加创意类型，详见枚举
	AdvancedCreativeType enum.AdvancedCreativeType `json:"advanced_creative_type,omitempty"`
	// Industry 广告主行业id，可以从获取行业接口进行获取
	Industry uint64 `json:"industry,omitempty"`
}

// Encode implement GetRequest interface
func (r ActionTextGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Add("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Add("landing_type", string(r.LandingType))
	if r.AdvancedCreativeType != "" {
		values.Add("advanced_creative_type", string(r.AdvancedCreativeType))
	}
	if r.Industry > 0 {
		values.Add("industry", strconv.FormatUint(r.Industry, 10))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// ActionTextGetResponse 行动号召字段内容获取 API Response
type ActionTextGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data []string `json:"data,omitempty"`
}
