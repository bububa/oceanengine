package aweme

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/aweme"
)

// EstimateProfit 获取随心推投放效果预估
// 小店获取投放效果预估
// 对于短视频预估播放量，对于直播预估观众数
func EstimateProfit(ctx context.Context, clt *core.SDKClient, accessToken string, req *aweme.EstimateProfitRequest) (*aweme.EstimateProfit, error) {
	var resp aweme.EstimateProfitResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/aweme/estimate_profit/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
