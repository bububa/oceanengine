package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// CarouselDelete 批量删除图集
// 通过此接口，用户可以对素材视频进行批量删除。当素材删除失败时，会展示在carousel_id列表，不在此列表内的素材表示删除成功！
func CarouselDelete(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.CarouselDeleteRequest) (*file.CarouselDeleteResult, error) {
	var resp file.CarouselDeleteResponse
	if err := clt.Post(ctx, "2/carousel/delete/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
