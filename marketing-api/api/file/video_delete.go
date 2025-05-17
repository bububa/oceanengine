package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// VideoDelete 批量删除视频素材
// 通过此接口，用户可以对素材视频进行批量删除。
// 当素材删除失败时，会展示在video_id列表，不在此列表内的素材表示删除成功！
func VideoDelete(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.VideoDeleteRequest) ([]string, error) {
	var resp file.VideoDeleteResponse
	err := clt.Post(ctx, "2/file/video/delete/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.FailVideoIDs, nil
}
