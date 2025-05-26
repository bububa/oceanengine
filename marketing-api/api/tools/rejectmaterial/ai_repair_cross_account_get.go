package rejectmaterial

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/rejectmaterial"
)

// AIRepairCrossAccountGet 根据mid查询同主体账户下修复建议列表
// 注意：本接口仅支持根据advertiser_id、拒审material_id，查询该拒审素材已被应用到哪些同主体账户下&已产生修复建议，注意：
// 1. 如传入的advertiser_id下不存在material_id  、material_id未产生修复建议，接口不会查询到数据
// 2. 该接口仅返回与传入advertiser_id同主体的账户
// 3. 通过应答参数获取到的related_advertiser_id、related_ai_repair_ids，可直接应用到采纳接口
func AIRepairCrossAccountGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *rejectmaterial.AIRepairCrossAccountGetRequest) (*rejectmaterial.AIRepairCrossAccountGetResult, error) {
	var resp rejectmaterial.AIRepairCrossAccountGetResponse
	if err := clt.GetAPI(ctx, "v3.0/reject_material/ai_repair/cross_account/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
