package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/local/file"
)

// VideoAwemeGet 获取素材库视频
func VideoAwemeGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.VideoAwemeGetRequest) (*file.VideoAwemeGetResult, error) {
	var resp file.VideoAwemeGetResponse
	if err := clt.GetAPI(ctx, "v3.0/local/file/video/aweme/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
