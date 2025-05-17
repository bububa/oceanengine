package shop

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/brand"
)

// AuthorizedGet 获取广告主绑定的品牌列表
func AuthorizedGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *brand.AuthorizedGetRequest) ([]brand.Brand, error) {
	var resp brand.AuthorizedGetResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/brand/authorized/get", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.BrandInfos, nil
}
