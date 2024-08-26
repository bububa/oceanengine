package agent

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/agent"
)

// Info 获取代理商信息
func Info(ctx context.Context, clt *core.SDKClient, accessToken string, req *agent.InfoRequest) ([]agent.Info, error) {
	var resp agent.InfoResponse
	if err := clt.Get(ctx, "2/agent/info/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
