package agent

import (
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// AdvertiserSelectRequest 代理商管理账户列表 API Request
type AdvertiserSelectRequest struct {
	// AdvertiserID 代理商ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Page 页码.默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面数据量.默认值: 100
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r AdvertiserSelectRequest) Encode() string {
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

// AdvertiserSelectResponse 代理商管理账户列表 API Response
type AdvertiserSelectResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *AdvertiserSelectResponseData `json:"data,omitempty"`
}

// AdvertiserSelectResponseData json返回值
type AdvertiserSelectResponseData struct {
	// List 广告主ID列表
	List []uint64 `json:"list,omitempty"`
	// AccountSource 账号列表的账号类型
	// 枚举值：AD广告主、STAR星图、LUBAN鲁班、DOMESTIC 入海
	AccountSource string `json:"account_source,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
