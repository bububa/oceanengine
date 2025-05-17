package rebate

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file/rebate"
)

// MaterialDownloadCreateTask 创建下载任务
// 根据筛选条件创建明点化素材相关指标的下载任务,返回用户query_id,用于后续的文件下载
func MaterialDownloadCreateTask(ctx context.Context, clt *core.SDKClient, accessToken string, req *rebate.MaterialDownloadCreateTaskRequest) (string, error) {
	var resp rebate.MaterialDownloadCreateTaskResponse
	if err := clt.Post(ctx, "2/file/rebate/material_download/create_task/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.QueryID, nil
}
