package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// VideoAd 上传视频
// 通过此接口，用户可以上传和广告相关的素材视频。
// 视频格式：mp4、mpeg、3gp、avi
func VideoAd(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.VideoAdRequest) (*file.Video, error) {
	var resp file.VideoAdResponse
	err := clt.Upload(ctx, "2/file/video/ad/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
