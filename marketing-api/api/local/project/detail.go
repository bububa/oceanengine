package project

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/local/project"
)

// Detail 获取项目详情
// 该接口暂不支持拉取推广目的为获取线索的项目
func Detail(ctx context.Context, clt *core.SDKClient, accessToken string, req *project.DetailRequest) (*project.ProjectDetail, error) {
	var resp project.DetailResponse
	if err := clt.GetAPI(ctx, "v3.0/local/project/detail/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
