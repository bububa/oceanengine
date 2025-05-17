package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// VideoAgentGet 代理商获取视频素材
func VideoAgentGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.VideoAgentGetRequest) (*file.VideoAgentGetResult, error) {
	var resp file.VideoAgentGetResponse
	if err := clt.GetAPI(ctx, "2/file/video/agent/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
