package task

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DetailRequest 获取星广联投(星图版)任务维度数据 API Request
type DetailRequest struct {
	// StarID 客户的星图id
	StarID uint64 `json:"star_id,omitempty"`
	// DemandID 任务id
	DemandID uint64 `json:"demand_id,omitempty"`
	// StarStartDate 查询起始时间，格式：yyyy-mm-dd ，只和安卓/iOS消耗、转化数、深度转化数相关
	StarStartDate string `json:"star_start_date,omitempty"`
	// StarEndDate 查询结束时间，格式：yyyy-mm-dd
	StarEndDate string `json:"star_end_date,omitempty"`
}

// Encode implement GetRequest interface
func (r DetailRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("star_id", strconv.FormatUint(r.StarID, 10))
	values.Set("demand_id", strconv.FormatUint(r.DemandID, 10))
	values.Set("star_start_date", r.StarStartDate)
	values.Set("star_end_date", r.StarEndDate)
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// DetailResponse 获取星广联投(星图版)任务维度数据 API Response
type DetailResponse struct {
	Data *Demand `json:"data,omitempty"`
	model.BaseResponse
}
