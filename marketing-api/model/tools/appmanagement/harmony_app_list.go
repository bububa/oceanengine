package appmanagement

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// HarmonyAppListRequest 查询鸿蒙应用列表 API Request
type HarmonyAppListRequest struct {
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
	// AccountAssetQueryScope 可选值:
	// ACCOCUNT
	// ALL
	// GROUP
	AccountAssetQueryScope string `json:"account_asset_query_scope,omitempty"`
}

// Encode implement GetRequest interface
func (r HarmonyAppListRequest) Encode() string {
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
	if r.AccountAssetQueryScope != "" {
		values.Set("account_asset_query_scope", r.AccountAssetQueryScope)
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}
