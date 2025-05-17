package ad

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/ad"
)

// SuggestBid 获取非ROI目标建议出价
// 支持通过该接口获取非ROI目标建议出价
func SuggestBid(ctx context.Context, clt *core.SDKClient, accessToken string, req *ad.SuggestBidRequest) (*ad.SuggestBidResult, error) {
	var resp ad.SuggestBidResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/suggest_bid/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
