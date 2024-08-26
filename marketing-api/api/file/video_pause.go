package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// VideoPause 按账户暂停素材
func VideoPause(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.VideoPauseRequest) (*file.VideoPauseResult, error) {
	var resp file.VideoPauseResponse
	if err := clt.Post(ctx, "2/file/video/pause/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
