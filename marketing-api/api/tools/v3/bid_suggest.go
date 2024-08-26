package v3

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools"
	v3 "github.com/bububa/oceanengine/marketing-api/model/tools/v3"
)

// BidSuggest 查询建议出价（巨量广告升级版）
// 通过广告分析查询广告的建议出价，目前仅支持手动投放的广告查询建议出价
func BidSuggest(ctx context.Context, clt *core.SDKClient, accessToken string, req *v3.BidSuggestRequest) (*tools.BidSuggest, error) {
	var resp v3.BidSuggestResponse
	if err := clt.GetAPI(ctx, "v3.0/tools/bids/suggest/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.Data, nil
}
