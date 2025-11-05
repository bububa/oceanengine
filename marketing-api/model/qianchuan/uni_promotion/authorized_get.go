package unipromotion

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/aweme"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AuthorizedGetRequest 获取可投全域推广抖音号列表 API Request
type AuthorizedGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 过滤器
	Filtering *AuthorizedGetFilter `json:"filtering,omitempty"`
	// Page 页码
	// 默认值： 1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小
	// 默认值：10，最大值：100
	PageSize int `json:"page_size,omitempty"`
}

type AuthorizedGetFilter struct {
	// MarketingGoal 营销目标，可选值:
	// LIVE_PROM_GOODS 直播全域（默认）
	// VIDEO_PROM_GOODS 商品全域
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// SearchKeyWords 检索词，支持根据抖音号show_Id/抖音号名搜索
	SearchKeyWords string `json:"search_key_words,omitempty"`
	// Scene 查询场景，可选值:
	// CREATE 计划创编页面（默认）
	// APPLY 申请授权页面
	Scene string `json:"scene,omitempty"`
}

// Encode implement GetRequest interface
func (r AuthorizedGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AuthorizedGetResponse 获取可投全域推广抖音号列表 API Response
type AuthorizedGetResponse struct {
	Data *AuthorizedGetResult `json:"data,omitempty"`
	model.BaseResponse
}

type AuthorizedGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// AwemeIDList 抖音号列表
	AwemeIDList []aweme.Aweme `json:"aweme_id_list,omitempty"`
}
