package advertiser

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/advertiser"
)

// DeliveryPkgSubmit 提交/编辑推广产品资质
// 用于提交以推广产品形式整组提交的投放资质，该接口可以同时用于新增和编辑。针对审核不通过的推广产品资质支持编辑提交
func DeliveryPkgSubmit(ctx context.Context, clt *core.SDKClient, accessToken string, req *advertiser.DeliveryPkgSubmitRequest) (*advertiser.DeliveryPkgSubmitResult, error) {
	var resp advertiser.DeliveryPkgSubmitResponse
	if err := clt.PostAPI(ctx, "v3.0/advertiser/delivery_pkg/submit/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
