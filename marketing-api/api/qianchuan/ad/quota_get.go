package ad

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/ad"
)

// QuotaGet 获取在投计划配额信息
// 此接口用于获取对应广告主的在投计划配额的相关信息，相关文档可见「在投计划配额」介绍文档
func QuotaGet(ctx context.Context, clt *core.SDKClient, accessToken string, advertiserID uint64) (*ad.QuotaGetResult, error) {
	var resp ad.QuotaGetResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/ad/quota/get/", &ad.QuotaGetRequest{
		AdvertiserID: advertiserID,
	}, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
