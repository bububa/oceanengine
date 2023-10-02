package tools

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// MicroAppListRequest 获取字节小程序 API Request
type MicroAppListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 过滤条件
	Filtering *MicroAppListFilter `json:"filtering,omitempty"`
	// Page 页码,默认值:1
	Page int `json:"page,omitempty"`
	// PageSize 分页大小,默认值:10，最大值100
	PageSize int `json:"page_size,omitempty"`
}

type MicroAppListFilter struct {
	// SearchType 搜索类型，可选值:
	// CREATE_ONLY只查询该账户创建的应用（默认值）
	// SHARE_ONLY只查询被共享的应用
	SearchType enum.MicroAppSearchType `json:"search_type,omitempty"`
	// SearchKey 小程序名称或备注的模糊匹配
	SearchKey string `json:"search_key,omitempty"`
	// AuditStatus 审核状态，可选值:
	// AUDIT_ACCEPTED 审核通过
	// AUDITING 审核中
	// AUDIT_REJECTED 审核不通过
	// ALL 全部（默认值）
	AuditStatus enum.MicroAppAuditStatus `json:"audit_status,omitempty"`
	// CreateTime 按创建时间查询的时间范围
	CreateTime *model.DateRange `json:"create_time,omitempty"`
}

// Encode implement GetRequest interface
func (r MicroAppListRequest) Encode() string {
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

// MicroAppListResponse 获取字节小程序 API Response
type MicroAppListResponse struct {
	model.BaseResponse
	Data *MicroAppListResult `json:"data,omitempty"`
}

type MicroAppListResult struct {
	// PageInfo 页面信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 字节小程序 列表
	List []MicroApp `json:"list,omitempty"`
}

// MicroApp 字节小程序
type MicroApp struct {
	// InstanceID 小程序资产id
	InstanceID uint64 `json:"instance_id,omitempty"`
	// Remark 字节小程序备注名称
	Remark string `json:"remark,omitempty"`
	// AppID 字节小程序app id
	AppID string `json:"app_id,omitempty"`
	// AuditStatus  审核状态:
	// AUDIT_ACCEPTED
	// AUDITING
	// AUDIT_REJECTED
	AuditStatus enum.MicroAppAuditStatus `json:"audit_status,omitempty"`
	// Reason 审核拒绝原因
	Reason string `json:"reason,omitempty"`
	// AdvertiserID 所属广告主账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CreateTime 创建时间
	CreateTime string `json:"create_time,omitempty"`
	// ModifyTime 修改时间
	ModifyTime string `json:"modify_time,omitempty"`
	// Category 所属账户类型
	Category uint64 `json:"category,omitempty"`
	// Name 字节小程序名称
	Name string `json:"name,omitempty"`
}
