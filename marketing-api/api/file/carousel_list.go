package file

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// CarouselList 获取图集素材
// 通过此接口，用户可以获取经过一定条件过滤后的广告主下创意素材库对应的图集及图集信息。
func CarouselList(clt *core.SDKClient, accessToken string, req *file.CarouselListRequest) (*file.CarouselListResult, error) {
	var resp file.CarouselListResponse
	if err := clt.Get("2/file/carousel/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
