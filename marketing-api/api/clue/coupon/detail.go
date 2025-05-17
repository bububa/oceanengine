package coupon

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/clue/coupon"
)

// Detail 获取卡券详情
func Detail(ctx context.Context, clt *core.SDKClient, accessToken string, req *coupon.DetailRequest) (*coupon.DetailResponseData, error) {
	var resp coupon.DetailResponse
	if err := clt.Get(ctx, "2/clue/coupon/detail/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
