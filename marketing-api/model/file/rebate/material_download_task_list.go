package rebate

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// QueryStatus 查询状态
type QueryStatus int

const (
	// QueryStatus_INIT 初始化
	QueryStatus_INIT QueryStatus = 1
	// QueryStatus_PROCESSING 运行中
	QueryStatus_PROCESSING QueryStatus = 2
	// QueryStatus_SUCCEED 成功
	QueryStatus_SUCCEED QueryStatus = 3
	// QueryStatus_FAILED 失败
	QueryStatus_FAILED QueryStatus = 4
)

// MaterialDownloadTaskListRequest 查询下载任务 API Request
type MaterialDownloadTaskListRequest struct {
	// AgentID 代理商帐户ID
	AgentID uint64 `json:"agent_id,omitempty"`
	// QueryID 下载任务查询ID（对应创建 下载任务接口 返回的query_id，建议必填）
	QueryID string `json:"query_id,omitempty"`
}

// Encode implements GetRequest interface
func (r MaterialDownloadTaskListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("agent_id", strconv.FormatUint(r.AgentID, 10))
	if r.QueryID != "" {
		values.Set("query_id", r.QueryID)
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// MaterialDownloadTaskListResponse 查询下载任务 API Response
type MaterialDownloadTaskListResponse struct {
	model.BaseResponse
	Data *MaterialDonwloadTaskListResult `json:"data,omitempty"`
}

type MaterialDonwloadTaskListResult struct {
	// QueryList 查询ID列表
	QueryList []MaterialDownloadTaskQuery `json:"query_list,omitempty"`
	// TaskList 任务ID列表
	TaskList []MaterialDownloadTask `json:"task_list,omitempty"`
	// QueryStatus 查询状态
	// 1- 初始化，2-运行中，3-成功，4-失败
	QueryStatus QueryStatus `json:"query_status,omitempty"`
}

// MaterialDownloadTaskQuery 查询ID
type MaterialDownloadTaskQuery struct {
	// QueryID 查询ID
	QueryID string `json:"query_id,omitempty"`
}

// MaterialDownloadTask 任务
type MaterialDownloadTask struct {
	// TaskID 任务ID
	TaskID string `json:"task_id,omitempty"`
	// CreateTime 任务创建时间
	CreateTime string `json:"create_time,omitempty"`
}
