package v3

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	v3 "github.com/bububa/oceanengine/marketing-api/model/v3"
)

// CdpBrandGet 获取关联云图的广告主账户信息
// 仅能查询到关联云图和cdp的广告主品牌及类别信息
func CdpBrandGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *v3.CdpBrandGetRequest) (*v3.CdpBrandGetResult, error) {
	var resp v3.CdpBrandGetResponse
	if err := clt.GetAPI(ctx, "v3.0/cdp/brand/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
