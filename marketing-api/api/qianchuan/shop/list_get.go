package shop

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/shop"
)

// AdvertiserList 获取店铺账户关联的广告账户列表
// 千川的广告账户无法通过接口直接获取，需要先通过【获取已授权的账户】获取客户授权时的店铺id/代理商id，再通过本接口或【获取代理商账户关联的广告账户列】获取广告账户列表
// 若需要知道返回的千川广告主账户是直投账户还是代理账户，请使用【获取千川广告账户全量信息】，通过返回账户的role进行判断
func AdvertiserList(ctx context.Context, clt *core.SDKClient, accessToken string, req *shop.AdvertiserListRequest) (*shop.AdvertiserListResponseData, error) {
	var resp shop.AdvertiserListResponse
	err := clt.Get(ctx, "v1.0/qianchuan/shop/advertiser/list/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
