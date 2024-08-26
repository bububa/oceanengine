package coupon

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/clue/coupon"
)

// Create 创建卡券
// 创建青鸟线索通卡券接口
func Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *coupon.CreateRequest) (uint64, error) {
	var resp coupon.CreateResponse
	if err := clt.Post(ctx, "2/clue/coupon/create/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.ActivityID, nil
}
