package file

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum/local"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// VideoUploadTaskListRequest 查询异步上传本地推视频结果 API Request
type VideoUploadTaskListRequest struct {
	// LocalAccountID 本地推账户id
	LocalAccountID uint64 `json:"local_account_id,omitempty"`
	// TaskIDs 任务id列表，单次最多支持100个任务id
	TaskIDs []uint64 `json:"task_ids,omitempty"`
}

// Encode implements GetRequest interface
func (r VideoUploadTaskListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("local_account_id", strconv.FormatUint(r.LocalAccountID, 10))
	values.Set("task_ids", string(util.JSONMarshal(r.TaskIDs)))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// VideoUploadTaskListResponse 查询异步上传本地推视频结果 API Response
type VideoUploadTaskListResponse struct {
	model.BaseResponse
	Data struct {
		// List 任务列表
		List []UploadTask `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

// UploadTask 上传任务
type UploadTask struct {
	// Status 任务处理状态 可选值:
	// PROCESS 处理中
	// SUCCESS 成功
	// FAILED 失败
	Status local.UploadTaskStatus `json:"status,omitempty"`
	// ErrorMsg 当任务失败后，会返回失败信息
	ErrorMsg string `json:"error_msg,omitempty"`
	// CreateTime 任务创建时间
	CreateTime string `json:"create_time,omitempty"`
	// TaskID 任务id
	TaskID uint64 `json:"task_id,omitempty"`
	// VideoInfo 视频信息
	VideoInfo *Video `json:"video_info,omitempty"`
}

func (t UploadTask) IsError() bool {
	return t.Status == local.UploadTaskStatus_FAILED
}

func (t UploadTask) Error() string {
	return util.StringsJoin("任务ID:", strconv.FormatUint(t.TaskID, 10), ", ", t.ErrorMsg)
}
