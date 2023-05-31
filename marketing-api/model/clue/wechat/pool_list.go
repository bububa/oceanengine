package wechat

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// PoolListRequest 获取微信库微信号列表 API Request
type PoolListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filter 过滤条件
	Filter *PoolListFilter `json:"filter,omitempty"`
	// Page 分页，页数>=1，默认值1
	Page int `json:"page,omitempty"`
	// PageSize 分页大小，1-100，默认值20
	PageSize int `json:"page_size,omitempty"`
}

// PoolListFilter 过滤条件
type PoolListFilter struct {
	// WechatNumber 微信号，模糊匹配
	WechatNumber string `json:"wechat_number,omitempty"`
	// NickName 微信昵称，模糊匹配
	NickName string `json:"nick_number,omitempty"`
	// InstanceID 微信号码包ID
	InstanceID uint64 `json:"instance_id,omitempty"`
	// AuditStatus 审核状态，允许值:
	// AUDITING: 审核中
	// AUDIT_REJECTED: 审核拒绝
	// AUDIT_ACCEPTED: 审核通过
	AuditStatus []AuditStatus `json:"audit_status,omitempty"`
}

// Encode implement GetRequest interface
func (r PoolListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filter != nil {
		values.Set("filter", string(util.JSONMarshal(r.Filter)))
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

// PoolListResponse 获取微信库微信号列表 API Response
type PoolListResponse struct {
	model.BaseResponse
	Data *PoolListData `json:"data,omitempty"`
}

type PoolListData struct {
	// Items 微信号列表
	Items []Wechat `json:"items,omitempty"`
	// TotalCount 总数
	TotalCount int64 `json:"total_count,omitempty"`
	// TotalPage 总页数
	TotalPage int `json:"total_page,omitempty"`
}
