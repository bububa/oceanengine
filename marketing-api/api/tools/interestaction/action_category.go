package interestaction

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/interestaction"
)

// ActionCategory 行为类目查询
func ActionCategory(ctx context.Context, clt *core.SDKClient, accessToken string, req *interestaction.ActionCategoryRequest) ([]interestaction.Object, error) {
	var resp interestaction.ActionCategoryResponse
	err := clt.Get(ctx, "2/tools/interest_action/action/category", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
