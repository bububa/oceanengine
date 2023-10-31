package dmp

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AudienceGroupGetRequest 获取人群分组 API Request
type AudienceGroupGetRequest struct {
	// AdvertiserID 千川广告主账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 过滤条件
	Filtering *AudienceGroupFilter `json:"filtering,omitempty"`
}

// AudienceGroupFilter 过滤条件
type AudienceGroupFilter struct {
	// Name 人群分组名称
	// 注意：模糊搜索
	Name string `json:"name,omitempty"`
}

// Encode implement GetRequest interface
func (r AudienceGroupGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AudienceGroupGetResponse 获取人群分组 API Response
type AudienceGroupGetResponse struct {
	model.BaseResponse
	Data struct {
		// AudienceGroupList 人群分组名称列表
		AudienceGroupList []string `json:"audience_group_list,omitempty"`
	} `json:"data,omitempty"`
}
