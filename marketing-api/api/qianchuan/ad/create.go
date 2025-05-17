package ad

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/ad"
)

// Create 创建计划（含创意生成规则）
// 用于创建一条完整的广告计划，包括投放规则（出价预算）、定向规则、创意（自定义创意/程序化创意）
// 与非千川的计划创建接口差异点在于，该接口将同时创建创意，即通过本接口可完成完整计划的创建
// 支持专业版和极速版计划的创建，具体传参要求参见字段描述
// 短视频带货目前暂不支持创建“计划托管”的广告计划
func Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *ad.CreateRequest) (*ad.CreateResult, error) {
	var resp ad.CreateResponse
	err := clt.Post(ctx, "v1.0/qianchuan/ad/create/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
