package clue

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// LifeGetRequest 获取本地推线索列表 API Request
type LifeGetRequest struct {
	// LocalAccountIDs 广告主ids，上限50
	LocalAccountIDs []uint64 `json:"local_account_ids,omitempty"`
	// StartTime 查询起始时间，格式：yyyy-MM-dd hh:mm:ss
	StartTime string `json:"start_time,omitempty"`
	// EndTime 查询截止时间，格式：yyyy-MM-dd hh:mm:ss
	EndTime string `json:"end_time,omitempty"`
	// Page 页码
	Page int `json:"page,omitempty"`
	// PageSize 页面大小, 默认值: 10，上限：100
	PageSize int `json:"page_size,omitempty"`
}

// Encode implements GetRequest interface
func (r LifeGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("local_account_ids", string(util.JSONMarshal(r.LocalAccountIDs)))
	values.Set("start_time", r.StartTime)
	values.Set("end_time", r.EndTime)
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

// LifeGetResponse 获取本地推线索列表 API Response
type LifeGetResponse struct {
	Data *LifeGetResult `json:"data,omitempty"`
	model.BaseResponse
}

type LifeGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 线索列表
	List []Clue `json:"list,omitempty"`
}
