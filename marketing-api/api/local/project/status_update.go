package project

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/local/project"
)

// StatusUpdate 批量更新项目状态
func StatusUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *project.StatusUpdateRequest) (*project.UpdateResult, error) {
	var resp project.StatusUpdateResponse
	if err := clt.PostAPI(ctx, "v3.0/local/project/status/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
