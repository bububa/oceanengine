package asynctask

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// CreateRequest 创建异步任务 API Request
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// TaskName 任务名称。最大长度 25 个字符，不能为空字符。
	TaskName string `json:"task_name,omitempty"`
	// Force true/false。是否强制生成新的任务（不复用之前任务的结果）。
	Force bool `json:"force,omitempty"`
	// TaskType 任务类型。可选值：“REPORT”（普通报表），"REPORT_DPA"（DPA 报表），"REPORT_BIDWORD"（关键词/搜索词报表）。
	TaskType enum.AsyncTaskType `json:"task_type,omitempty"`
	// TaskParams 任务的参数
	TaskParams *TaskParams `json:"task_params,omitempty"`
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	bs, _ := json.Marshal(r)
	return bs
}

// CreateResponse 创建异步任务 API Response
type CreateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *Task `json:"data,omitempty"`
}
