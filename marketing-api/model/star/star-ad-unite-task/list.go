package task

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ListRequest 获取星广联投(星图版)任务列表 API Request
type ListRequest struct {
	// Filtering 过滤器
	Filtering *ListFilter `json:"filtering,omitempty"`
	// StarID 客户的星图id
	StarID uint64 `json:"star_id,omitempty"`
	// Page 页码，默认为1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认10，最大值20
	PageSize int `json:"page_size,omitempty"`
}

type ListFilter struct {
	// Status 任务状态 可选值:
	// ALL 不限
	// CANCELED 已取消
	// FINISHED 已完成
	// ONGOING 进行中
	// RECEIVEING 待接收
	// WAIT_EVALUATE 待评价
	// WAIT_PAYMENT 待付款
	Status enum.StarAdUniteTaskStatus `json:"status,omitempty"`
}

// Encode implement GetRequest interface
func (r ListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("star_id", strconv.FormatUint(r.StarID, 10))
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// ListResponse 获取星广联投(星图版)任务列表 API Response
type ListResponse struct {
	Data *ListResult `json:"data,omitempty"`
	model.BaseResponse
}

type ListResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// Demands 任务信息
	Demands []Demand `json:"demands,omitempty"`
}
