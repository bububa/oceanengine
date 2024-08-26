package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// VideoMaterialClearTaskResultGet 下载清理任务结果
// 根据adv_id和clear_id返回低效/同质视频素材的清理结果，与「创建素材清理任务」、「获取清理任务列表」接口配合使用
func VideoMaterialClearTaskResultGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.VideoMaterialClearTaskResultGetRequest) (*file.VideoMaterialClearTaskResultGetData, error) {
	var resp file.VideoMaterialClearTaskResultGetResponse
	if err := clt.Get(ctx, "2/file/video/material/clear_task_result/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
