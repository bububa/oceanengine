package advertiser

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/advertiser"
)

// DeliveryPkgGet 查询推广产品资质
// 用于查询广告主以推广产品形式提交的投放资质，可以获取到资质审核状态等信息
func DeliveryPkgGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *advertiser.DeliveryPkgGetRequest) (*advertiser.DeliveryPkg, error) {
	var resp advertiser.DeliveryPkgGetResponse
	if err := clt.GetAPI(ctx, "v3.0/advertiser/delivery_pkg/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.DeliveryPkg, nil
}
