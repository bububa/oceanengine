package project

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/project"
)

// BudgetUpdate 更新项目预算
func BudgetUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *project.BudgetUpdateRequest) (*project.UpdateResponseData, error) {
	var resp project.UpdateResponse
	if err := clt.PostAPI(ctx, "v3.0/project/budget/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
