package agent

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/agent"
)

// QueryRiskPromotionList 【代理商】查询广告违规信息
// 通过此接口，用户可以获取代理商账户下：
// 支持获取在投放中图片、视频和落地页被拒审的巨量广告信息，仅展示广告拒审时的信息
// 支持获取广告中未过审的素材信息以及这个素材还在同代理商的哪些广告下（只披露近7天有消耗的关联广告）
func QueryRiskPromotionList(ctx context.Context, clt *core.SDKClient, accessToken string, req *agent.QueryRiskPromotionListRequest) (*agent.QueryRiskPromotionListResult, error) {
	var resp agent.QueryRiskPromotionListResponse
	if err := clt.GetAPI(ctx, "2/agent/query/risk_promotion_list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
