package v3

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/report/asynctask"
	v3 "github.com/bububa/oceanengine/marketing-api/model/report/asynctask/v3"
)

// Create 自定义报表—创建异步任务
// 自定义报表—创建异步任务。每个开发者每天最多只能为每个用户创建 10 个任务（不包括提交失败的任务）。
func Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *v3.CreateRequest) (*asynctask.Task, error) {
	var resp asynctask.CreateResponse
	if err := clt.PostAPI(ctx, "v3.0/report/custom/async_task/create/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
