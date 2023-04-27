package file

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// VideoMaterialClearTaskGetRequest 获取清理任务列表 API Request
type VideoMaterialClearTaskGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ClearIDs 清理任务id列表，最多支持10个
	ClearIDs []uint64 `json:"clear_ids,omitempty"`
	// Page 页码，默认：1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认10，范围：1-100
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r VideoMaterialClearTaskGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if len(r.ClearIDs) > 0 {
		values.Set("clear_ids", string(util.JSONMarshal(r.ClearIDs)))
	}
	if r.Page > 1 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// VideoMaterialClearTaskGetResponse 获取清理任务列表 API Response
type VideoMaterialClearTaskGetResponse struct {
	model.BaseResponse
	Data *VideoMaterialClearTaskGetData `json:"data,omitempty"`
}

type VideoMaterialClearTaskGetData struct {
	List     []VideoMaterialClearTask `json:"list,omitempty"`
	PageInfo *model.PageInfo          `json:"page_info,omitempty"`
}

// VideoMaterialClearTask 清理任务
type VideoMaterialClearTask struct {
	// ClearID 清理任务id
	ClearID uint64 `json:"clear_id,omitempty"`
	// ClearTaskParams 任务的参数
	ClearTaskParams *VideoMaterialClearTaskParameters `json:"clear_task_params,omitempty"`
	// CreateTime 任务创建时间，格式：%Y-%m-%d %H:%M:%S
	CreateTime string `json:"create_time,omitempty"`
	// TaskStatus 任务状态:
	// DONE 已完成
	// RUNNING 运行中
	// TIMEOUT超时
	TaskStatus enum.VideoMaterialClearTaskStatus `json:"task_status,omitempty"`
}
