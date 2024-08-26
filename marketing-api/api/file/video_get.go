package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// VideoGet 获取视频素材
// 通过此接口，用户可以获取经过一定条件过滤后的广告主下创意素材库对应的视频及视频信息。
func VideoGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.VideoGetRequest) (*file.VideoGetResponseData, error) {
	var resp file.VideoGetResponse
	err := clt.Get(ctx, "2/file/video/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
