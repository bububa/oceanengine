package agent

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/agent"
)

// ChildAgentSelect 二级代理商列表
// 获取代理商下的广告主ID列表，如果存在二级代理，要获取二级代理商列表请使用agent/child_agent/select/接口。如果您需要获取二级代理商下的广告主，需要传递二级代理商id调用此接口即可获得。如果要查看广告主ID的详细信息请参考广告主信息接口。
func ChildAgentSelect(ctx context.Context, clt *core.SDKClient, accessToken string, req *agent.ChildAgentSelectRequest) ([]uint64, error) {
	var resp agent.ChildAgentSelectResponse
	if err := clt.Get(ctx, "2/agent/child_agent/select/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.ChildAgentIDs, nil
}
