package abtest

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/abtest"
)

// Delete 删除实验
// 仅实验状态为进行中的实验，可以删除
func Delete(ctx context.Context, clt *core.SDKClient, accessToken string, req *abtest.DeleteRequest) (*abtest.DeleteResult, error) {
	var resp abtest.DeleteResponse
	if err := clt.Post(ctx, "2/tools/ab_test/delete/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
