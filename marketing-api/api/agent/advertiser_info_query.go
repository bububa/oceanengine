package agent

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/agent"
)

// AdvertiserInfoQuery 广告主账户信息查询
// 广告主账户信息查询，包括：基础信息、自运营报备标签、账户绑定时间、账户联系人、账户负责人和协作者等
func AdvertiserInfoQuery(ctx context.Context, clt *core.SDKClient, accessToken string, req *agent.AdvertiserInfoQueryRequest) ([]agent.AccountDetail, error) {
	var resp agent.AdvertiserInfoQueryResponse
	if err := clt.Get(ctx, "2/agent/advertiser_info/query/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.AccountDetailList, nil
}
