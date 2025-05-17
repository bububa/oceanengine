package ad

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/ad"
)

// Create 创建广告计划
// 计划包括了计划名称、投放范围、投放目标、用户定向、预算与出价，对于其中的概念解释可以参考：【广告计划】
// 计划创建涉及了多个资产管理：应用管理、落地页管理、转化目标管理、定向包管理、DMP人群管理等；开发者需要提前根据开放接口构建这些资产功能，以免创建计划卡住
// 新版营销链路广告计划差异点
// 新版计划中去除投放内容中落地页信息（旧版计划中在计划维度）
// 新版计划中新增广告位信息（旧版计划中在创意维度）
// 新版计划中增加监测链接内容（旧版计划中在创意维度）
func Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *ad.CreateRequest) (uint64, error) {
	var resp ad.CreateResponse
	err := clt.Post(ctx, "2/ad/create/", req, &resp, accessToken)
	if err != nil {
		return 0, err
	}
	return resp.Data.AdID, nil
}
