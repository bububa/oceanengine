package coupon

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/clue/coupon"
)

// CodeConsume 核销券码
func CodeConsume(ctx context.Context, clt *core.SDKClient, accessToken string, req *coupon.CodeConsumeRequest) error {
	return clt.Post(ctx, "2/clue/coupon/code/consume/", req, nil, accessToken)
}
