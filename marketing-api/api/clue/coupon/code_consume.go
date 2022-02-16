package coupon

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/clue/coupon"
)

// CodeConsume 核销券码
func CodeConsume(clt *core.SDKClient, accessToken string, req *coupon.CodeConsumeRequest) error {
	return clt.Post("2/clue/coupon/code/consume/", req, nil, accessToken)
}
