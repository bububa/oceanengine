package task

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ItemListRequest 获取星广联投(星图版)视频维度数据 API Request
type ItemListRequest struct {
	// StarID 客户的星图id
	StarID uint64 `json:"star_id,omitempty"`
	// Page 页码，默认为1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认10，最大值20
	PageSize int `json:"page_size,omitempty"`
	// DemandID 任务id
	DemandID uint64 `json:"demand_id,omitempty"`
	// StarStartDate 查询起始时间，格式：yyyy-mm-dd ，只和安卓/iOS消耗、转化数、深度转化数相关
	StarStartDate string `json:"star_start_date,omitempty"`
	// StarEndDate 查询结束时间，格式：yyyy-mm-dd
	StarEndDate string `json:"star_end_date,omitempty"`
}

// Encode implement GetRequest interface
func (r ItemListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("star_id", strconv.FormatUint(r.StarID, 10))
	values.Set("demand_id", strconv.FormatUint(r.DemandID, 10))
	values.Set("star_start_date", r.StarStartDate)
	values.Set("star_end_date", r.StarEndDate)
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

// ItemListResponse 获取星广联投(星图版)视频维度数据 API Response
type ItemListResponse struct {
	model.BaseResponse
	Data struct {
		StatInfo []ItemStatInfo `json:"stat_info,omitempty"`
	} `json:"data,omitempty"`
}
