package finance

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/finance"
)

// DetailGet 获取账户钱包信息
// 此接口用于获取账户钱包内余额信息，支持千川广告主/代理商帐户
func DetailGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *finance.DetailGetRequest) (*finance.DetailGetResult, error) {
	var resp finance.DetailGetResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/finance/detail/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
