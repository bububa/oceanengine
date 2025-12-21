package unipromotion

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	unipromotion "github.com/bububa/oceanengine/marketing-api/model/qianchuan/uni_promotion"
)

// AdRoi2GoalUpdate 更新全域推广控成本计划支付ROI目标
// 支持更新商品&直播全域推广控成本计划支付ROI目标
// 注意：仅支持优化目标为“整体支付roi”的计划，不支持“净成交支付roi”
func AdRoi2GoalUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *unipromotion.AdRoi2GoalUpdateRequest) ([]unipromotion.AdUpdateResult, error) {
	var resp unipromotion.AdUpdateResponse
	if err := clt.Post(ctx, "v1.0/qianchuan/uni_promotion/ad/roi2_goal/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.Results, nil
}
