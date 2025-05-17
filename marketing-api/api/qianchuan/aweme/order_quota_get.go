package aweme

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/aweme"
)

// OrderQuotaGet 查询随心推使用中订单配额信息
// 查询小店随心推使用中订单配额信息，更多信息可参考：【小店随心推订单配额介绍文档】
func OrderQuotaGet(ctx context.Context, clt *core.SDKClient, accessToken string, advertiserID uint64) (*aweme.OrderQuota, error) {
	var resp aweme.OrderQuotaGetResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/aweme/order/quota/get/", &aweme.OrderQuotaGetRequest{AdvertiserID: advertiserID}, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
