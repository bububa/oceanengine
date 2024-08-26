package coupon

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/clue/coupon"
)

// CodeGet 查询券码记录
func CodeGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *coupon.CodeGetRequest) (*coupon.CodeGetResponseData, error) {
	var resp coupon.CodeGetResponse
	if err := clt.Get(ctx, "2/clue/coupon/code/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
