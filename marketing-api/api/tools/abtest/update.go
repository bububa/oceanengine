package abtest

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/abtest"
)

// Update 更改实验
func Update(ctx context.Context, clt *core.SDKClient, accessToken string, req *abtest.UpdateRequest) (uint64, error) {
	var resp abtest.UpdateResponse
	if err := clt.Post(ctx, "2/tools/ab_test/update/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.TestID, nil
}
