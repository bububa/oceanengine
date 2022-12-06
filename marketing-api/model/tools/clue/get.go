package clue

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetRequest 获取线索列表 API Request
type GetRequest struct {
	// AdvertiserIDs 广告主ids
	AdvertiserIDs []string `json:"advertiser_ids,omitempty"`
	// StartTime 查询起始时间，格式：yyyy-MM-dd 或者 yyyy-MM-dd hh:mm:ss
	StartTime string `json:"start_time,omitempty"`
	// EndTime 查询截止时间，格式：yyyy-MM-dd 或者 yyyy-MM-dd hh:mm:ss
	EndTime string `json:"end_time,omitempty"`
	// Page 页数 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小 默认值: 10
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r GetRequest) Encode() string {
	values := util.GetUrlValues()
	ids, _ := json.Marshal(r.AdvertiserIDs)
	values.Set("advertiser_ids", string(ids))
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

// GetResponse 获取线索列表 API Response
type GetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *GetResponseData `json:"data,omitempty"`
}

// GetResponseData json返回值
type GetResponseData struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 线索列表
	List []Clue `json:"list,omitempty"`
}
