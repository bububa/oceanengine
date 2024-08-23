package blueflow

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// KeywordListRequest 获取广告下可用蓝海关键词 API Request
type KeywordListRequest struct {
	// Filtering 过滤器
	Filtering *KeywordListFilter `json:"filtering,omitempty"`
	// AdvertiserID 广告账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ProjectID 项目ID
	ProjectID uint64 `json:"project_id,omitempty"`
}

type KeywordListFilter struct {
	// BlueFlowKeywordName 蓝海关键词名称，模糊匹配，长度限制50字
	BlueFlowKeywordName string `json:"blue_flow_keyword_name,omitempty"`
	// BlueFlowRecommendType 关键词推荐类别，可选值:
	// ORDINARY 代表普通
	// PREFERRED 代表优选
	BlueFlowRecommendType string `json:"blue_flow_recommend_type,omitempty"`
}

// Encode implements GetRequest interface
func (r KeywordListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("project_id", strconv.FormatUint(r.ProjectID, 10))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// KeywordListResponse 获取广告下可用蓝海关键词 API Response
type KeywordListResponse struct {
	model.BaseResponse
	Data struct {
		// List 蓝海流量包关键词列表
		List []Keyword `json:"list,omitempty"`
	} `json:"data,omitempty"`
}
