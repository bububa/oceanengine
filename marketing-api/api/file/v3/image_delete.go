package v3

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	v3 "github.com/bububa/oceanengine/marketing-api/model/file/v3"
)

// ImageDelete 批量删除图片素材
// 通过此接口，用户可以批量删除广告主下创意素材库的图片。
func ImageDelete(ctx context.Context, clt *core.SDKClient, accessToken string, req *v3.ImageDeleteRequest) ([]string, error) {
	var resp v3.ImageDeleteResponse
	err := clt.PostAPI(ctx, "v3.0/file/image/delete/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.FailImageIDs, nil
}
