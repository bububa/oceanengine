package businessplatform

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// PartnerOrganizationListRequest 查询合作组织 API Request
type PartnerOrganizationListRequest struct {
	// OrganizationID 纵横组织id
	OrganizationID uint64 `json:"organization_id,omitempty"`
	// Page 页码. 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面数据量. 默认值: 10
	PageSize int `json:"page_size,omitempty"`
	// Filtering 过滤条件
	Filtering *PartnerOrganizationListFilter `json:"filtering,omitempty"`
}

// PartnerOrganizationListFilter 过滤条件
type PartnerOrganizationListFilter struct {
	// PartnerOrganizationIDs 合作的巨量纵横组织id，最大不超过200
	PartnerOrganizationIDs []uint64 `json:"partner_organization_ids,omitempty"`
	// Status 合作状态，默认查询绑定成功的状态
	// 允许值：BOUND（已绑定）、BINDING（绑定中）、INVALID（失效）、DELETED（删除）
	Status enum.PartnerOrganizationStatus `json:"status,omitempty"`
}

// Encode implement GetRequest interface
func (r PartnerOrganizationListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("organization_id", strconv.FormatUint(r.OrganizationID, 10))
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

// PartnerOrganizationListResponse 查询合作组织 API Response
type PartnerOrganizationListResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *PartnerOrganizationListData `json:"data,omitempty"`
}

// PartnerOrganizationListData json返回值
type PartnerOrganizationListData struct {
	// List 列表
	List []PartnerOrganization `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
