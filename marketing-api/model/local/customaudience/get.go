package customaudience

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum/local"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetRequest 查询本地推创编可用人群包 API Request
type GetRequest struct {
	// LocalAccountID 本地推广告主ID
	LocalAccountID uint64 `json:"local_account_id,omitempty"`
	// TagsType 按人群包属性筛选，允许值：
	// CUSTOM 自定义人群包
	// SYS_RECOMMEND 系统推荐人群包
	TagsType local.CustomAudienceTagsType `json:"tags_type,omitempty"`
	// Page 页码，默认值1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值20，最大值100
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements GetRequest interface
func (r GetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("local_account_id", strconv.FormatUint(r.LocalAccountID, 10))
	if r.TagsType != "" {
		values.Set("tags_type", string(r.TagsType))
	}
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

// GetResponse 查询本地推创编可用人群包 API Response
type GetResponse struct {
	model.BaseResponse
	Data *GetResult `json:"data,omitempty"`
}

type GetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// CustomAudienceList 人群包列表
	CustomAudienceList []CustomAudience `json:"custom_audience_list,omitempty"`
}
