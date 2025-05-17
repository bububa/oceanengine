package asynctask

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/report/asynctask"
)

// Get 获取任务列表
func Get(ctx context.Context, clt *core.SDKClient, accessToken string, req *asynctask.GetRequest) (*asynctask.GetResponseData, error) {
	var resp asynctask.GetResponse
	err := clt.Get(ctx, "2/async_task/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
