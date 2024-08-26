package asynctask

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/report/asynctask"
)

// Create 创建异步任务
// 创建异步任务，每个开发者每天最多只能为每个用户创建 10 个任务（不包括提交失败的任务）。
// 当前支持两种报表类型：普通报表和 DPA 报表。支持的分组条件及类型详见【附录：分组条件】
// 普通报表的分组条件组合规则：ID类型 * [时间类型] * [细分类型]
// DPA 报表的分组条件组合规则：DPA 类型 * [ID 类型] * [STAT_GROUP_BY_TIME_DAY|STAT_GROUP_BY_TIME_HOUR] * [STAT_GROUP_BY_INVENTORY|STAT_GROUP_BY_EXTERNAL_ACTION|(STAT_GROUP_BY_INVENTORY * STAT_GROUP_BY_EXTERNAL_ACTION)|STAT_GROUP_BY_IMAGE_MODE|STAT_GROUP_BY_PLATFORM]
// 关键词报表的分组条件组合规则：STAT_GROUP_BY_BIDWORD_ID * [STAT_GROUP_BY_AD_ID] * [时间类型] * [STAT_GROUP_BY_INVENTORY|STAT_GROUP_BY_PRICING]
// 搜索词报表的分组条件组合规则：STAT_GROUP_BY_BIDWORD_ID * STAT_GROUP_BY_AD_ID * STAT_GROUP_BY_QUERY * [STAT_GROUP_BY_TIME_DAY｜STAT_GROUP_BY_TIME_WEEK] * [STAT_GROUP_BY_IMAGE_MODE|STAT_GROUP_BY_PRICING]
// 暂不支持获取当天的数据
// 暂不支持查询小时纬度分组数据（DPA报表支持）
func Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *asynctask.CreateRequest) (*asynctask.Task, error) {
	var resp asynctask.CreateResponse
	if err := clt.Post(ctx, "2/async_task/create/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
