package dmp

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dmp"
)

// BrandGet 获取广告账户关联云图账户信息
// 根据真实的客户adv_id 获取有权限的云图品牌列表信息，信息包括：品牌id、品牌名称、虚拟advid
func BrandGet(ctx context.Context, clt *core.SDKClient, accessToken string, advertiserID uint64) ([]dmp.BrandInfo, error) {
	req := dmp.BrandGetRequest{
		AdvertiserID: advertiserID,
	}
	var resp dmp.BrandGetResponse
	if err := clt.Get(ctx, "2/dmp/brand/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.BrandInfo, nil
}
