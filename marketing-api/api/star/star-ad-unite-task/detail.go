package task

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	task "github.com/bububa/oceanengine/marketing-api/model/star/star-ad-unite-task"
)

// Detail 获取星广联投(星图版)任务维度数据
// 此接口可根据星图客户id和星广联投(星图版)任务id，获取某段时间下该任务信息和投放数据。
func Detail(ctx context.Context, clt *core.SDKClient, accessToken string, req *task.DetailRequest) (*task.Demand, error) {
	var resp task.DetailResponse
	err := clt.GetAPI(ctx, "2/star/star_ad_unite_task/detail/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
