package project

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/local/project"
)

// Create 创建项目
// 创建 巨量引擎体验版 项目 广告主下可创建项目数：200
// 项目创建成功后，默认为启用状态
func Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *project.CreateRequest) (uint64, error) {
	var resp project.CreateResponse
	if err := clt.PostAPI(ctx, "v3.0/local/project/create/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.ProjectID, nil
}
