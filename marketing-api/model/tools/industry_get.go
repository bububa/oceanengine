package tools

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// IndustryGetRequest 获取行业列表 API Request
type IndustryGetRequest struct {
	// Level 只获取某级别数据，1:第一级,2:第二级,3:第三级，默认都返回
	Level int `json:"level,omitempty"`
	// Type 可选值："ADVERTISER"，"AGENT"，"ADVERTISER"为原有广告3.0行业, "AGENT"为代理商行业获取，代理商行业level都为1
	Type string `json:"type,omitempty"`
}

// Encode implement GetRequest interface
func (r IndustryGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("level", strconv.Itoa(r.Level))
	values.Set("type", r.Type)
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// IndustryGetResponse 获取行业列表 API Response
type IndustryGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		List []Industry `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

// Industry 行业
type Industry struct {
	// IndustryID 行业ID
	IndustryID uint64 `json:"industry_id,omitempty"`
	// IndustryName 名称
	IndustryName string `json:"industry_name,omitempty"`
	// Level 所在级别，1：一级行业、2：二级行业、3：三级行业
	Level int `json:"level,omitempty"`
	// FirstIndustryID 该行业的一级行业ID
	FirstIndustryID uint64 `json:"first_industry_id,omitempty"`
	// FirstIndustryName 该行业的一级行业名称
	FirstIndustryName string `json:"first_industry_name,omitempty"`
	// SecondIndustryID 该行业的二级行业ID
	SecondIndustryID uint64 `json:"second_industry_id,omitempty"`
	// SecondIndustryName 该行业的二级行业名称
	SecondIndustryName string `json:"second_industry_name,omitempty"`
	// ThirdIndustryID 该行业的三级行业ID
	ThirdIndustryID uint64 `json:"third_industry_id,omitempty"`
	// ThirdIndustryName 该行业的三级行业名称
	ThirdIndustryName string `json:"third_industry_name,omitempty"`
}
