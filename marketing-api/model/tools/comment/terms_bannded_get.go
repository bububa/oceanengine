package comment

import (
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// TermsBanndedGetRequest 获取屏蔽词列表 API Request
type TermsBanndedGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Page 页数 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小默认值: 10
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r TermsBanndedGetRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	return values.Encode()
}

// TermsBanndedGetResponse 获取屏蔽词列表 API Response
type TermsBanndedGetResponse struct {
	model.BaseResponse
	// Data json 返回值
	Data *TermsBanndedGetResponseData `json:"data,omitempty"`
}

// TermsBanndedGetResponseData json 返回值
type TermsBanndedGetResponseData struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// Terms 屏蔽词列表
	Terms []string `json:"terms,omitempty"`
}
