package rejectmaterial

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/rejectmaterial"
)

// AIRepairAcceptTaskCreate 创建采纳「拒审素材修复建议」任务
// 注意：本接口为异步接口，需与「获取素材修复建议采纳结果」接口搭配使用
// 一次传入50个修复id，从创建采纳任务→采纳完成，时效预估在1分钟以内
func AIRepairAcceptTaskCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *rejectmaterial.AIRepairAcceptTaskCreateRequest) (*rejectmaterial.AIRepairAcceptTaskCreateResult, error) {
	var resp rejectmaterial.AIRepairAcceptTaskCreateResponse
	if err := clt.PostAPI(ctx, "v3.0/reject_material/ai_repair_accept_task/create/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
