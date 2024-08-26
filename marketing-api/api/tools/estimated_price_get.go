package tools

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools"
)

// EstimatedPriceGet 获取预估点击成本
func EstimatedPriceGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *tools.EstimatedPriceGetRequest) (*tools.EstimatedPrice, error) {
	var resp tools.EstimatedPriceGetResponse
	err := clt.Get(ctx, "2/tools/estimated_price/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
