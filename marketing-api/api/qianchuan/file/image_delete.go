package file

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/file"
)

// ImageDelete 批量删除图片素材
// 通过此接口，用户可以批量删除广告主下创意素材库的图片。
func ImageDelete(clt *core.SDKClient, accessToken string, req *file.ImageDeleteRequest) ([]string, error) {
	var resp file.ImageDeleteResponse
	if err := clt.Post("v1.0/qianchuan/file/image/delete/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.FailImageIDs, nil
}
