package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// AudioAd 上传图文内的音频素材
func AudioAd(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.AudioAdRequest) (*file.Audio, error) {
	var resp file.AudioAdResponse
	err := clt.Upload(ctx, "2/file/audio/ad/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.AudioInfo, nil
}
