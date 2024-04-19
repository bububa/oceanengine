package customercenter

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AdvertiserListRequest 获取纵横组织下资产账户列表（分页） API Request
type AdvertiserListRequest struct {
	// CcAccountID 纵横组织id，通过【获取已授权账户】接口获取
	CcAccountID uint64 `json:"cc_account_id,omitempty"`
	// Page 页码. 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面数据量. 默认值: 10
	PageSize int `json:"page_size,omitempty"`
	// Filtering 过滤条件
	Filtering *AdvertiserListFilter `json:"filtering,omitempty"`
	// AccountSource 账户类型，可选值：
	// AD 广告主账号、默认
	// ENTERPRISE企业号
	// LOCAL：本地推
	AccountSource string `json:"account_source"`
}

// AdvertiserListFilter 过滤条件
type AdvertiserListFilter struct {
	// AccountName 根据账户名称过滤
	AccountName string `json:"account_name,omitempty"`
}

// Encode implement GetRequest interface
func (r AdvertiserListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("cc_account_id", strconv.FormatUint(r.CcAccountID, 10))
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
	if r.AccountSource != "" {
		values.Set("account_source", r.AccountSource)
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AdvertiserListResponse 获取纵横组织下资产账户列表（分页）API Response
type AdvertiserListResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *AdvertiserListData `json:"data,omitempty"`
}

// AdvertiserListData json返回值
type AdvertiserListData struct {
	// List 账户列表
	List []Advertiser `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
