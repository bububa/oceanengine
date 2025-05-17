package tools

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools"
)

// SearchBidRatioGet 获取快投推荐出价系数
func SearchBidRatioGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *tools.SearchBidRatioGetRequest) (float64, error) {
	var resp tools.SearchBidRatioGetResponse
	if err := clt.GetAPI(ctx, "2/tools/search_bid_ratio/get/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.Ratio, nil
}
