package majordomo

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/majordomo"
)

// AdvertiserSelect 获取纵横组织下资产账户列表
// 获取客户中心管家下的广告主ID列表。注意：如果发现没有查询到一些广告主，需要您在客户中心确认是否这些广告主是否和此管家解绑关系。
func AdvertiserSelect(ctx context.Context, clt *core.SDKClient, accessToken string, advertiserID uint64) ([]majordomo.Advertiser, error) {
	req := &majordomo.AdvertiserSelectRequest{
		AdvertiserID: advertiserID,
	}
	var resp majordomo.AdvertiserSelectResponse
	if err := clt.Get(ctx, "2/majordomo/advertiser/select", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
