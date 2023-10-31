package eventmanager

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ShareGetRequest 事件管理资产查看共享范围 API Request
type ShareGetRequest struct {
	// OrganizationID 纵横组织id
	OrganizationID uint64 `json:"organization_id,omitempty"`
	// AssetID 事件管理资产id
	AssetID uint64 `json:"asset_id,omitempty"`
	// Page 页码，默认值为1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值为10，最大不超过200
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r ShareGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("organization_id", strconv.FormatUint(r.OrganizationID, 10))
	values.Set("asset_id", strconv.FormatUint(r.AssetID, 10))
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

// ShareGetResponse 事件管理资产查看共享范围 API Response
type ShareGetResponse struct {
	model.BaseResponse
	// Data 返回数据
	Data *ShareGetData `json:"data,omitempty"`
}

// ShareGetData .
type ShareGetData struct {
	// List 共享账户ID（adv+bpid+枚举值）集合
	List []ShareInfo `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
