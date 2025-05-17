package rebate

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// MaterialDownloadFileRequest 下载任务结果
type MaterialDownloadFileRequest struct {
	// AgentID 代理商帐户ID
	AgentID uint64 `json:"agent_id,omitempty"`
	// TaskID 任务ID
	TaskID string `json:"task_id,omitempty"`
}

// Encode implements GetRequest interface
func (r MaterialDownloadFileRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("agent_id", strconv.FormatUint(r.AgentID, 10))
	values.Set("task_id", r.TaskID)
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// MaterialDownloadFileResponse 下载任务结果
type MaterialDownloadFileResponse struct {
	model.BaseResponse
	Data json.RawMessage `json:"data,omitempty"`
}
