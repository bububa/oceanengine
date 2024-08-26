package ad

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/ad"
)

// Update 更新计划（含创意生成规则）
// 用于更新一条完整的广告计划，包括投放规则（出价预算），定向规则，创意（自定义创意/程序化创意）
// 与非千川的计划更新接口差异点在于，该接口将同时更新创意，即通过本接口可完成完整计划的更新
// 目前的更新方式为全量更新，可先通过【获取计划详情】接口获取计划详情内容后，再进行更新
// 支持专业推广和极速推广计划的更新，具体可参看字段表
func Update(ctx context.Context, clt *core.SDKClient, accessToken string, req *ad.UpdateRequest) (*ad.CreateResult, error) {
	var resp ad.CreateResponse
	err := clt.Post(ctx, "v1.0/qianchuan/ad/update/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
