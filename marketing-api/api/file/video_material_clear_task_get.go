package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// VideoMaterialClearTaskGet 获取清理任务列表
// 返回已创建的低效/同质视频素材清理任务列表
func VideoMaterialClearTaskGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.VideoMaterialClearTaskGetRequest) (*file.VideoMaterialClearTaskGetData, error) {
	var resp file.VideoMaterialClearTaskGetResponse
	if err := clt.Get(ctx, "2/file/video/material/clear_task/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
