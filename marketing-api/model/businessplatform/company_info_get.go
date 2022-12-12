package businessplatform

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CompanyInfoGetRequest 获取纵横组织下所有主体信息 API Request
type CompanyInfoGetRequest struct {
	// OrganizationID 纵横组织id
	OrganizationID uint64 `json:"organization_id,omitempty"`
	// Page 页码. 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面数据量. 默认值: 10
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r CompanyInfoGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("organization_id", strconv.FormatUint(r.OrganizationID, 10))
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

// CompanyInfoGetResponse 获取纵横组织下所有主体信息 API Response
type CompanyInfoGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *CompanyInfoGetData `json:"data,omitempty"`
}

// CompanyInfoGetData json返回值
type CompanyInfoGetData struct {
	// List 主体信息列表
	List []CompanyInfo `json:"company_info,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
