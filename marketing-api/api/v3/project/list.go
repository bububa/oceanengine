package project

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/project"
)

// List 获取广告项目列表
func List(ctx context.Context, clt *core.SDKClient, accessToken string, req *project.ListRequest) (*project.ListResponseData, error) {
	var resp project.ListResponse
	if err := clt.GetAPI(ctx, "v3.0/project/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
