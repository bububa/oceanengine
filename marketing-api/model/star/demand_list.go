package star

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DemandListRequest 获取星图客户任务列表 API Request
type DemandListRequest struct {
	// StarID 星图id，星图客户授权后，通过“获取已授权账户”接口，查询到账号角色为”6-星图账号“的账户id，即为星图id
	StarID uint64 `json:"star_id,omitempty"`
	// Filtering 过滤条件，若此字段不传，或传空则视为无限制条件
	Filtering *DemandListFilter `json:"filtering,omitempty"`
	// Page 页码，默认为1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认10，最大值50
	PageSize int `json:"page_size,omitempty"`
}

// DemandListFilter 过滤条件
type DemandListFilter struct {
	// ComponentType 组件类型
	ComponentType enum.StarComponentType `json:"component_type,omitempty"`
	// Name 任务名称，模糊匹配
	Name string `json:"name,omitempty"`
	// UniversalOrderStatus 任务订单状态（订单状态，并非任务状态，意为过滤出包含该状态订单的任务）
	UniversalOrderStatus enum.UniversalOrderStatus `json:"universal_order_status,omitempty"`
	// UniversalSettlementType 结算方式
	UniversalSettlementType enum.UniversalSettlementType `json:"universal_settlement_type,omitempty"`
}

// Encode implement GetRequest interface
func (r DemandListRequest) Encode() string {
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

// DemandListResponse 获取星图客户任务列表 API Response
type DemandListResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *DemandListResponseData `json:"data,omitempty"`
}

// DemandListResponseData json返回值
type DemandListResponseData struct {
	List     []Demand       `json:"list,omitempty"`
	PageInfo model.PageInfo `json:"page_info"`
}
