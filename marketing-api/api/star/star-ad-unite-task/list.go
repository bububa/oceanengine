package task

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	task "github.com/bububa/oceanengine/marketing-api/model/star/star-ad-unite-task"
)

// List 获取星广联投(星图版)任务列表
// 此接口可根据星图客户id，获取星图客户账号下星广联投(星图版)全部任务。
func List(ctx context.Context, clt *core.SDKClient, accessToken string, req *task.ListRequest) (*task.ListResult, error) {
	var resp task.ListResponse
	err := clt.GetAPI(ctx, "2/star/star_ad_unite_task/list/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
