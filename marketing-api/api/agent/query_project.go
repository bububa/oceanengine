package agent

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/agent"
)

// QueryProject 查询项目信息
func QueryProject(ctx context.Context, clt *core.SDKClient, accessToken string, req *agent.QueryProjectRequest) (*agent.QueryProjectResult, error) {
	var resp agent.QueryProjectResponse
	if err := clt.GetAPI(ctx, "2/query/project/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
