package appmanagement

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/appmanagement"
)

// UploadTaskCreate 创建异步文件上传任务
// 执行异步上传操作「支持所有账户体系」
func UploadTaskCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *appmanagement.UploadTaskCreateRequest) (uint64, error) {
	var resp appmanagement.UploadTaskCreateResponse
	if err := clt.Post(ctx, "2/tools/app_management/upload_task/create/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.UploadID, nil
}
