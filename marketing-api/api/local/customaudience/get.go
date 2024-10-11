package customaudience

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/local/customaudience"
)

// Get 查询本地推创编可用人群包
func Get(ctx context.Context, clt *core.SDKClient, accessToken string, req *customaudience.GetRequest) (*customaudience.GetResult, error) {
	var resp customaudience.GetResponse
	if err := clt.GetAPI(ctx, "v3.0/local/custom_audience/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
