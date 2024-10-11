package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/local/file"
)

// VideoUpload 上传视频
func VideoUpload(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.VideoUploadRequest) (*file.Video, error) {
	var resp file.VideoUploadResponse
	if err := clt.Upload(ctx, "v3.0/local/file/video/upload/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
