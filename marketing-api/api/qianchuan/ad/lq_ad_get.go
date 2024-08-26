package ad

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/ad"
)

// LqAdGet 获取低效计划列表
// 用于获取千川广告账户下的低效计划列表
// 注意：
// 低效计划一定是在投计划；关停或删除后，将不再是低效计划
// 创意重复度高或低活跃的计划，都有可能是低效计划。具体定义待平台后续公布
// 请勿使用定时任务高频轮询；建议实现“低效计划暂停/清理”功能，由前端触发调用该接口
func LqAdGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *ad.LqAdGetRequest) ([]uint64, error) {
	var resp ad.LqAdGetResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/lq_ad/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.AdIDs, nil
}
