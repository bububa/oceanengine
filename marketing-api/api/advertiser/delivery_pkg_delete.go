package advertiser

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/advertiser"
)

// DeliveryPkgDelete 批量删除推广产品资质
// 该接口支持批量删除审核拒绝的投放资质（推广产品）。
// 删除失败的原因包括2类：① 传入了审核状态≠「审核不通过」的资质id；② 广告主账户下没有该资质。
func DeliveryPkgDelete(ctx context.Context, clt *core.SDKClient, accessToken string, req *advertiser.DeliveryPkgDeleteRequest) (*advertiser.DeliveryPkgDeleteResult, error) {
	var resp advertiser.DeliveryPkgDeleteResponse
	if err := clt.PostAPI(ctx, "v3.0/advertiser/delivery_pkg/delete/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
