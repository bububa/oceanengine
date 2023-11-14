package advertiser

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DeliveryPkgConfigRequest 查询推广产品资质规则配置 API Request
type DeliveryPkgConfigRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// FirstIndustryID 一级行业id
	FirstIndustryID uint64 `json:"first_industry_id,omitempty"`
	// SecondIndustryID 二级行业id
	SecondIndustryID uint64 `json:"second_industry_id,omitempty"`
	// ThirdIndustryID 三级行业id
	ThirdIndustryID uint64 `json:"third_industry_id,omitempty"`
}

// Encode implement GetRequest interface
func (r DeliveryPkgConfigRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("first_industry_id", strconv.FormatUint(r.FirstIndustryID, 10))
	values.Set("second_industry_id", strconv.FormatUint(r.SecondIndustryID, 10))
	values.Set("third_industry_id", strconv.FormatUint(r.ThirdIndustryID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// DeliveryPkgConfigResponse 查询推广产品资质规则配置 API Response
type DeliveryPkgConfigResponse struct {
	model.BaseResponse
	Data struct {
		// IndustryConfig 资质规则
		IndustryConfig *IndustryConfig `json:"industry_config,omitempty"`
	} `json:"data,omitempty"`
}
