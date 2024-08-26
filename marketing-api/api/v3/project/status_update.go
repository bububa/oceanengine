package project

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/project"
)

// StatusUpdate 更新项目状态
func StatusUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *project.StatusUpdateRequest) (*project.UpdateResponseData, error) {
	var resp project.UpdateResponse
	if err := clt.PostAPI(ctx, "v3.0/project/status/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
