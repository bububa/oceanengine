package rejectmaterial

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/rejectmaterial"
)

// AIRepairAcceptTaskList 获取采纳素材修复建议任务结果
// 注意：本接口为异步接口，需与创建采纳「拒审素材修复建议」任务API搭配使用
// 您可通过本接口查询采纳的最新状态
func AIRepairAcceptTaskList(ctx context.Context, clt *core.SDKClient, accessToken string, req *rejectmaterial.AIRepairAcceptTaskListRequest) ([]rejectmaterial.AIRepairAcceptTask, error) {
	var resp rejectmaterial.AIRepairAcceptTaskListResponse
	if err := clt.GetAPI(ctx, "v3.0/reject_material/ai_repair_accept_task/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.ResultList, nil
}
