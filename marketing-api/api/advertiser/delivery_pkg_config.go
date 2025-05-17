package advertiser

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/advertiser"
)

// DeliveryPkgConfig 查询推广产品资质规则配置
// 根据商业化行业获取不同行业下对应的资质提交规则。需注意：每个行业下的资质提交规则可能会因平台及外部监管的要求而发生变化，当规则发生变更时，规则的版本号+1
func DeliveryPkgConfig(ctx context.Context, clt *core.SDKClient, accessToken string, req *advertiser.DeliveryPkgConfigRequest) (*advertiser.IndustryConfig, error) {
	var resp advertiser.DeliveryPkgConfigResponse
	if err := clt.GetAPI(ctx, "v3.0/advertiser/delivery_pkg_config/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.IndustryConfig, nil
}
