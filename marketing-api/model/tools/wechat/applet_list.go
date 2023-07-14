package wechat

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AppletListRequest 获取微信小程序列表 API Request
type AppletListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 过滤条件
	Filtering *AppletListFilter `json:"filtering,omitempty"`
	// Page 页码，默认值为1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值为10，最大值100
	PageSize int `json:"page_size,omitempty"`
}

// AppletListFilter 过滤条件
type AppletListFilter struct {
	// Name 小程序名称或备注的模糊匹配
	Name string `json:"name,omitempty"`
	// AuditStatus 审核状态 允许值: AUDIT_ACCEPTED 审核通过、AUDITING 审核中、AUDIT_REJECTED 审核不通过。不传表示全部状态
	AuditStatus enum.WechatAuditStatus `json:"audit_status,omitempty"`
	// SearchType  搜索类型
	// 允许值： CREATE_ONLY 只查询该账户创建的应用（默认值）、SHARE_ONLY 只查询被共享的应用
	SearchType enum.WechatSearchType `json:"search_type,omitempty"`
	// CreateTime 按创建时间查询的时间范围
	CreateTime *ListTimeRange `json:"create_time,omitempty"`
}

// Encode implement GetRequest interface
func (r AppletListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
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

// AppletListResponse 获取微信小程序列表 API Response
type AppletListResponse struct {
	model.BaseResponse
	Data *AppletListResult `json:"data,omitempty"`
}

type AppletListResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 微信小程序列表
	List []WechatApplet `json:"list,omitempty"`
}
