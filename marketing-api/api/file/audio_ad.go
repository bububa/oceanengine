package file

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// AudioAd 上传图文内的音频素材
// 通过此接口，用户可以上传和广告相关的音频图片，例如图文中的音频。
func AudioAd(clt *core.SDKClient, accessToken string, req *file.AudioAdRequest) (*file.Audio, error) {
	var resp file.AudioAdResponse
	if err := clt.Upload("2/file/video/ad/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.AudioInfo, nil
}
