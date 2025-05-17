package ad

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/ad"
)

// EstimateEffect 获取预估效果接口
// 推直播间-日常销售-通投-托管-放量投放场景获取预估效果
func EstimateEffect(ctx context.Context, clt *core.SDKClient, accessToken string, req *ad.EstimateEffectRequest) (*ad.EstimateEffectResult, error) {
	var resp ad.EstimateEffectResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/estimate/effect/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
