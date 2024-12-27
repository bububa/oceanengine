package agent

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/agent"
)

// QueryStatement 查询项目关联结算单信息
// 代理商查询项目关联结算单信息
func QueryStatement(ctx context.Context, clt *core.SDKClient, accessToken string, req *agent.QueryStatementRequest) ([]agent.ProjectRefStatement, error) {
	var resp agent.QueryStatementResponse
	if err := clt.GetAPI(ctx, "2/query/statement/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.ProjectRefStatmentList, nil
}
