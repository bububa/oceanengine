package asynctask

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetRequest 获取任务列表 API Request
type GetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 筛选条件
	Filtering *GetFilter `json:"filtering,omitempty"`
	// Page 分搜索页码
	// 默认值: 1
	// 取值范围: ≥ 0
	Page int `json:"page,omitempty"`
	// PageSize 一页展示的数据数量
	// 默认值: 1
	// 取值范围: 1-10
	PageSize int `json:"page_size,omtiempty"`
}

// GetFilter 筛选条件
type GetFilter struct {
	// TaskIDs 任务 id。最多 10 个
	TaskIDs []uint64 `json:"task_ids,omitempty"`
	// TaskName 任务名称。
	TaskName string `json:"task_name,omitempty"`
}

// Encode implement GetRequest interface
func (r GetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		bs, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(bs))
	}
	if r.Page > 1 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 1 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// GetResponse 获取任务列表 API Response
type GetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *GetResponseData `json:"date,omitempty"`
}

// GetResponseData json返回值
type GetResponseData struct {
	PageInfo model.PageInfo `json:"page_info"`
	List     []Task         `json:"list,omitempty"`
}
