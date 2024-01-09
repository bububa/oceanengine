package file

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/file"
)

// CarouselAdGet 获取同主体下广告主图集素材
// 通过此接口，用户可以查询获取同主体下的广告主图集素材信息。
func CarouselAdGet(clt *core.SDKClient, accessToken string, req *file.CarouselAdGetRequest) (*file.CarouselAdGetResult, error) {
	var resp file.CarouselAdGetResponse
	if err := clt.Get("2/carousel/ad/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
