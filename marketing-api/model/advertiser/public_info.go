package advertiser

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// PublicInfoRequest 广告主公开信息 API Request
type PublicInfoRequest struct {
	// AdvertiserIDs 广告主ID集合(如果包含没有访问权限的ID,将返回no permission error) 取值范围: 1-100
	AdvertiserIDs []uint64 `json:"advertiser_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r PublicInfoRequest) Encode() string {
	idsBytes, _ := json.Marshal(r.AdvertiserIDs)
	values := util.GetUrlValues()
	values.Set("advertiser_ids", string(idsBytes))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// PublicInfoResponse 广告主公开信息 API Response
type PublicInfoResponse struct {
	model.BaseResponse
	// Data json返回值
	Data []PublicInfo `json:"data,omitempty"`
}

// PublicInfo 广告主公开信息
type PublicInfo struct {
	// ID 广告主ID
	ID uint64 `json:"id,omitempty"`
	// Name 账户名
	Name string `json:"name,omitempty"`
	// Company 公司名
	Company string `json:"company,omitempty"`
	// FirstIndustryName 一级行业名
	FirstIndustryName string `json:"first_industry_name,omitempty"`
	// SecondIndustryName 二级行业名
	SecondIndustryName string `json:"second_industry_name,omitempty"`
}
