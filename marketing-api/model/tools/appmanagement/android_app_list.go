package appmanagement

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AndroidAppListRequest 查询安卓应用信息（支持所有账户体系）API Request
type AndroidAppListRequest struct {
	// AccountID 账户id，accout_type类型对应账户ID
	AccountID uint64 `json:"account_id,omitempty"`
	// AccountType 账户类型
	// 允许值：BP 巨量纵横组织、AD 广告主账号、 STAR 星图
	AccountType enum.AccountType
	// Filtering 过滤条件
	Filtering *AppListFilter `json:"filtering,omitempty"`
	// Page 页码，默认值为1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值为10，最大不超过200
	PageSize int `json:"page_size,omitempty"`
}

// AppListFilter 过滤条件
type AppListFilter struct {
	// SearchKey 搜索关键字
	// appid或者应用名，可以为空，可以传中文
	// 长度不超过50
	SearchKey string `json:"search_key,omitempty"`
	// SearchType 搜索类型:
	// ALL:查询全部，包括创建和被共享的应用（默认）
	// CREATE_ONLY:只查询广告主创建的应用
	// SHARED_ONLY:只查询被共享的应用
	SearchType AppSearchType `json:"search_type,omitempty"`
	// Status 应用状态:
	// ALL:所有状态
	// AUDIT_DOING:审核中
	// AUDIT_REJECTED:审核失败
	// BOOKING:预约中
	// ENABLE:已发布（默认）
	// PAST_DUE:已逾期
	Status AppStatus `json:"status,omitempty"`
	// CreateTime 按创建时间查询的时间范围
	CreateTime *TimeRange `json:"create_time,omitempty"`
	// PublishTime 按发布时间查询的时间范围
	PublishTime *TimeRange `json:"publish_time,omitempty"`
	// ScheduledPublishTime 按预约发布时间查询的时间范围
	ScheduledPublishTime *TimeRange `json:"scheduled_publish_time,omitempty"`
}

// Encode implement GetRequest interface
func (r AndroidAppListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AccountID, 10))
	values.Set("account_type", string(r.AccountType))
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	if r.Page != 1 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}
