package aweme

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum/local"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AuthorizedGetRequest 获取本地推创编可用抖音号 API Request
type AuthorizedGetRequest struct {
	// LocalAccountID 本地推广告主ID
	LocalAccountID uint64 `json:"local_account_id,omitempty"`
	// Filtering 过滤器
	Filtering *AuthorizedGetFilter `json:"filtering,omitempty"`
	// Page 页码
	Page int `json:"page,omitempty"`
	// PageSize 页面大小
	PageSize int `json:"page_size,omitempty"`
	// MarketingGoal 抖音号使用场景，允许值
	// LIVE 直播
	// VIDEO_IMAGE 短视频/图文
	MarketingGoal local.MarketingGoal `json:"marketing_goal,omitempty"`
}

type AuthorizedGetFilter struct {
	// SearchKeyWord 根据抖音号id和名称进行搜索
	SearchKeyWord string `json:"search_key_word,omitempty"`
}

// Encode implements GetRequest interface
func (r AuthorizedGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("local_account_id", strconv.FormatUint(r.LocalAccountID, 10))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	if r.MarketingGoal != "" {
		values.Set("marketing_goal", string(r.MarketingGoal))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AuthorizedGetResponse 获取本地推创编可用抖音号 API Response
type AuthorizedGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *AuthorizedGetResult `json:"data,omitempty"`
}

type AuthorizedGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// AwemeIDList 抖音号列表
	AwemeIDList []Aweme `json:"aweme_id_list,omitempty"`
}
