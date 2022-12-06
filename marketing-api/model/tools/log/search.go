package log

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// SearchRequest 日志查询 API Request
type SearchRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ObjectID 操作对象ID, 1 <= len <= 20 , 可以为campaign_id、ad_id、creative_id， 各种id可以随意组合
	ObjectID []uint64 `json:"object_id,omitempty"`
	// StartTime 日志查询开始时间，格式 "2019-07-24 21:46:57"
	StartTime string `json:"start_time,omitempty"`
	// EndTime 日志查询结束时间，格式 "2019-07-24 21:46:57"
	EndTime string `json:"end_time,omitempty"`
	// Page 页码
	Page int `json:"page,omitempty"`
	// PageSize 页面大小
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r SearchRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if len(r.ObjectID) > 0 {
		ids, _ := json.Marshal(r.ObjectID)
		values.Set("object_id", string(ids))
	}
	if r.StartTime != "" {
		values.Set("start_time", r.StartTime)
	}
	if r.EndTime != "" {
		values.Set("end_time", r.EndTime)
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

// SearchResponse 日志查询 API Response
type SearchResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *SearchResponseData `json:"data,omitempty"`
}

// SearchResponseData json返回值
type SearchResponseData struct {
	// Logs 日志详情
	Logs []Log `json:"logs,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
