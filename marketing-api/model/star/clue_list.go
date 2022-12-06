package star

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ClueListRequest 获取星图订单投后线索 API Request
type ClueListRequest struct {
	// StarID 星图id，星图客户授权后，通过“获取已授权账户”接口，查询到账号角色为”6-星图账号“的账户id，即为星图id
	StarID uint64 `json:"star_id,omitempty"`
	// DemandID 任务id
	DemandID uint64 `json:"demand_id,omitempty"`
	// OrderID 订单id
	OrderID uint64 `json:"order_id,omitempty"`
	// Page 页码，默认为1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认10，最大值50
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r ClueListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("star_id", strconv.FormatUint(r.StarID, 10))
	if r.DemandID > 0 {
		values.Set("demand_id", strconv.FormatUint(r.DemandID, 10))
	}
	if r.OrderID > 0 {
		values.Set("order_id", strconv.FormatUint(r.OrderID, 10))
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

// ClueListResponse 获取星图客户任务列表 API Response
type ClueListResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *ClueListResponseData `json:"data,omitempty"`
}

// ClueListResponseData json返回值
type ClueListResponseData struct {
	List     []Clue         `json:"list,omitempty"`
	PageInfo model.PageInfo `json:"page_info"`
}
