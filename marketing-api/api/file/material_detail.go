package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// MaterialDetail 查询素材标签信息
// 素材信息查询，根据提供的素材 id 查询素材属性信息，目前仅支持视频素材
func MaterialDetail(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.MaterialDetailRequest) ([]file.Material, error) {
	var resp file.MaterialDetailResponse
	if err := clt.Get(ctx, "2/file/material/detail/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.Materials, nil
}
