package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/local/file"
)

// VideoGet 获取素材库视频
func VideoGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.VideoGetRequest) (*file.VideoGetResult, error) {
	var resp file.VideoGetResponse
	if err := clt.GetAPI(ctx, "v3.0/local/file/video/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
