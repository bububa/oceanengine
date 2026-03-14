package ebp

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// https://open.oceanengine.com/labels/7/docs/1829550825614739?origin=left_nav

// AdvertiserSelectRequest 获取升级版巨量引擎工作台下账户列表 API Request
type AdvertiserListRequest struct {
	// EnterpriseOrganizationID 升级版巨量引擎工作台ID【必填】
	EnterpriseOrganizationID uint64 `json:"enterprise_organization_id"`
	// AccountSource 账户类型【必填】
	// 枚举值：AD 巨量营销客户账号、LOCAL 本地推账户
	AccountSource string `json:"account_source"`
	// Filtering 过滤器
	Filtering *AdvertiserListFilter `json:"filtering,omitempty"`
	// Page 页码.默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面数据量[1,100].默认值: 10
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r AdvertiserListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("enterprise_organization_id", strconv.FormatUint(r.EnterpriseOrganizationID, 10))
	values.Set("account_source", r.AccountSource)
	if r.Filtering != nil {
		bs, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(bs))
	}
	if r.Page > 1 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

type AdvertiserListFilter struct {
	// AccountName 广告主名称
	AccountName string `json:"account_name,omitempty"`
	// ActiveAccount 是否是活跃账户
	ActiveAccount bool `json:"active_account,omitempty"`
}

// AdvertiserListResponse 获取升级版巨量引擎工作台下账户列表 API Response
type AdvertiserListResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *AdvertiserListResponseData `json:"data,omitempty"`
}

// AdvertiserListResponseData json返回值
type AdvertiserListResponseData struct {
	// AccountList 广告主账户列表
	AccountList []AdvertiserListResponseAccount `json:"account_list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

type AdvertiserListResponseAccount struct {
	// AccountID 账户id
	AccountID uint64 `json:"account_id,omitempty"`
	// AccountType 账户类型
	// 枚举值：AD_NORMAL 普通客户账号、DOU_PLUSDOU+ 类客户账号、LOCAL 本地推客户账号
	AccountType string `json:"account_type,omitempty"`
	// AccountName 账户名称
	AccountName string `json:"account_name,omitempty"`
}
