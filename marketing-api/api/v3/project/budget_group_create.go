package project

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/project"
)

// BudgetGroupCreate  创建预算组
// 新建预算组，预算组可设置多个项目的预算（日预算）
func BudgetGroupCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *project.BudgetGroupCreateRequest) (uint64, error) {
	var resp project.BudgetGroupCreateResponse
	if err := clt.PostAPI(ctx, "v3.0/budget_group/create/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.BudgetGroupID, nil
}
