package coupon

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/clue/coupon"
)

// Get 获取卡券列表
func Get(ctx context.Context, clt *core.SDKClient, accessToken string, req *coupon.GetRequest) (*coupon.GetResponseData, error) {
	var resp coupon.GetResponse
	if err := clt.Get(ctx, "2/clue/coupon/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
