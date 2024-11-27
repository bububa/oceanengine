package star

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DemandOrderListRequest 获取星图客户任务列表  API Request
type DemandOrderListRequest struct {
	// Filtering 过滤条件，若此字段不传，或传空则视为无限制条件
	Filtering *DemandOrderListFilter `json:"filtering,omitempty"`
	// StarID 星图id，星图客户授权后，通过“获取已授权账户”接口，查询到账号角色为”6-星图账号“的账户id，即为星图id
	StarID uint64 `json:"star_id,omitempty"`
	// DemandID 任务id，可通过“获取星图客户任务列表”获取
	DemandID uint64 `json:"demand_id,omitempty"`
	// Page 页码，默认为1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认10，最大值50
	PageSize int `json:"page_size,omitempty"`
}

// DemandOrderListFilter 过滤条件
type DemandOrderListFilter struct {
	// UniversalOrderStatus 任务订单状态（订单状态，并非任务状态，意为过滤出包含该状态订单的任务）
	UniversalOrderStatus enum.UniversalOrderStatus `json:"universal_order_status,omitempty"`
}

// Encode implement GetRequest interface
func (r DemandOrderListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("star_id", strconv.FormatUint(r.StarID, 10))
	if r.Filtering != nil {
		buf, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(buf))
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

// DemandOrderListResponse 获取星图客户任务列表  API Response
type DemandOrderListResponse struct {
	// Data json返回值
	Data *DemandOrderListResult `json:"data,omitempty"`
	model.BaseResponse
}

// DemandOrderListResult json返回值
type DemandOrderListResult struct {
	List     []Order        `json:"list,omitempty"`
	PageInfo model.PageInfo `json:"page_info"`
}
