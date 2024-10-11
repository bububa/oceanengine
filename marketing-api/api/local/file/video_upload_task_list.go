package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/local/file"
)

// VideoUploadTaskList 查询异步上传本地推视频结果
func VideoUploadTaskList(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.VideoUploadTaskListRequest) ([]file.UploadTask, error) {
	var resp file.VideoUploadTaskListResponse
	if err := clt.GetAPI(ctx, "v3.0/local/file/video/upload_task/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
