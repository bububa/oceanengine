package abtest

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/abtest"
)

// Stop 关停实验
// 实验状态非进行中的实验，可以关停。
func Stop(ctx context.Context, clt *core.SDKClient, accessToken string, req *abtest.StopRequest) (*abtest.StopResult, error) {
	var resp abtest.StopResponse
	if err := clt.Post(ctx, "2/tools/ab_test/stop/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
