package project

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/project"
)

// Create 创建项目
func Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *project.CreateRequest) (*project.CreateResult, error) {
	var resp project.CreateResponse
	if err := clt.PostAPI(ctx, "v3.0/project/create/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
