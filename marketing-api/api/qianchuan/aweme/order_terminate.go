package aweme

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/aweme"
)

// OrderTerminate 终止随心推订单
// 该接口用于终止小店随心推订单么，终止成功后将会返回当前广告主终止订单配额信息，订单终止条件以及配额相关信息具体见：「小店随心推」订单配额介绍文档(对外)
func OrderTerminate(ctx context.Context, clt *core.SDKClient, accessToken string, req *aweme.OrderTerminateRequest) (*aweme.OrderTerminateResult, error) {
	var resp aweme.OrderTerminateResponse
	if err := clt.Post(ctx, "v1.0/qianchuan/aweme/order/terminate/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
