package appmanagement

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/appmanagement"
)

// UploadTaskList 查询文件异步上传任务
// 查询异步上传解析任务的状态信息「支持所有账户体系」
func UploadTaskList(clt *core.SDKClient, accessToken string, req *appmanagement.UploadTaskListRequest) ([]appmanagement.UploadTask, error) {
	var resp appmanagement.UploadTaskListResponse
	if err := clt.Get("2/tools/app_management/upload_task/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
