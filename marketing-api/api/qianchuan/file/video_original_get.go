package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/file"
)

// VideoOriginalGet 获取首发素材
// 支持查询素材是否是首发素材，传入素材ID列表，返回首发素材列表。
func VideoOriginalGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.VideoOriginalGetRequest) ([]string, error) {
	var resp file.VideoOriginalGetResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/file/video/original/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.OriginalMaterialIDs, nil
}
