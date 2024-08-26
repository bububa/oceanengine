package project

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/project"
)

// CostProtectStatusGet 批量获取项目成本保障状态
// 本接口支持批量查询项目的成本保障状态及相关赔付信息（接口能力仅对齐巨量广告升级版），对应平台能力如下图所示
func CostProtectStatusGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *project.CostProtectStatusGetRequest) (*project.CostProtectStatusGetResult, error) {
	var resp project.CostProtectStatusGetResponse
	if err := clt.GetAPI(ctx, "v3.0/project/cost_protect_status/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
