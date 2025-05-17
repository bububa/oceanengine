package aweme

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/aweme"
)

// AuthorizedGet 获取千川账户下已授权抖音号
// 获取千川账户下已授权抖音号
func AuthorizedGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *aweme.AuthorizedGetRequest) (*aweme.AuthorizedGetResponseData, error) {
	var resp aweme.AuthorizedGetResponse
	err := clt.Get(ctx, "v1.0/qianchuan/aweme/authorized/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
