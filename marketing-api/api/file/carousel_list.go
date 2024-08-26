package file

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// CarouselList 获取图集素材
// 通过此接口，用户可以获取经过一定条件过滤后的广告主下创意素材库对应的图集及图集信息。
func CarouselList(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.CarouselListRequest) (*file.CarouselListResult, error) {
	var resp file.CarouselListResponse
	if err := clt.Get(ctx, "2/carousel/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
