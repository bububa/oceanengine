package project

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/local/project"
)

// List 获取项目列表
// 注意：该接口暂不支持拉取推广目的为获取线索的项目
func List(ctx context.Context, clt *core.SDKClient, accessToken string, req *project.ListRequest) (*project.ListResult, error) {
	var resp project.ListResponse
	if err := clt.GetAPI(ctx, "v3.0/local/project/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
