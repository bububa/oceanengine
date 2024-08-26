package rebate

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file/rebate"
)

// MaterialDownloadTaskList 查询下载任务
// 查询指定query_id的所有下载任务
func MaterialDownloadTaskList(ctx context.Context, clt *core.SDKClient, accessToken string, req *rebate.MaterialDownloadTaskListRequest) (*rebate.MaterialDonwloadTaskListResult, error) {
	var resp rebate.MaterialDownloadTaskListResponse
	if err := clt.Get(ctx, "2/file/rebate/material_download/get_download_task_list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
