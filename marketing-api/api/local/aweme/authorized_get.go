package aweme

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/local/aweme"
)

// AuthorizedGet 获取本地推创编可用抖音号
func AuthorizedGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *aweme.AuthorizedGetRequest) (*aweme.AuthorizedGetResult, error) {
	var resp aweme.AuthorizedGetResponse
	if err := clt.GetAPI(ctx, "v3.0/local/aweme/authorized/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
