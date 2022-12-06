package businessplatform

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CompanyAccountGetRequest 获取主体下的账户列表 API Request
type CompanyAccountGetRequest struct {
	// OrganizationID 纵横组织id
	OrganizationID uint64 `json:"organization_id,omitempty"`
	// CompanyID 公司主体id
	CompanyID uint64 `json:"company_id,omitempty"`
	// AccountType 账户类型，可选值:
	// AD: 广告账户
	// QIANCHUAN:千川广告账户
	AccountType []enum.AccountType `json:"account_type,omitempty"`
	// Page 页码. 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面数据量. 默认值: 10
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r CompanyAccountGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("organization_id", strconv.FormatUint(r.OrganizationID, 10))
	values.Set("company_id", strconv.FormatUint(r.CompanyID, 10))
	bs, _ := json.Marshal(r.AccountType)
	values.Set("account_type", string(bs))
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

// CompanyAccountGetResponse 获取主体下的账户列表 API Response
type CompanyAccountGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *CompanyAccountGetData `json:"data,omitempty"`
}

// CompanyAccountGetData json返回值
type CompanyAccountGetData struct {
	// List 账户列表
	List []Account `json:"account_list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
