package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// VideoMaterialClearTaskCreate 创建素材清理任务
// 创建低效/同质素材清理任务的异步接口，最多同时创建10个运行中的清理任务，配合「获取清理任务列表」、「下载清理任务结果」接口使用
func VideoMaterialClearTaskCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.VideoMaterialClearTaskCreateRequest) (uint64, error) {
	var resp file.VideoMaterialClearTaskCreateResponse
	if err := clt.Post(ctx, "2/file/video/material/clear_task/create/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.ClearID, nil
}
