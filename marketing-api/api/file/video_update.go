package file

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// VideoUpdate 更新视频
// 通过此接口，用户可以批量更新素材视频的名称。
func VideoUpdate(clt *core.SDKClient, accessToken string, req *file.VideoUpdateRequest) ([]file.VideoForUpdate, error) {
	var resp file.VideoUpdateResponse
	err := clt.Post("2/file/video/update/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	if resp.Data == nil {
		return nil, nil
	}
	return resp.Data.Videos, nil
}
