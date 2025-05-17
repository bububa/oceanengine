package analyse

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/product/analyse"
)

// CompareCreative 商品竞争分析详情-创意比对
// 获取商品竞争分析详情-创意比对详情
func CompareCreative(ctx context.Context, clt *core.SDKClient, accessToken string, req *analyse.CompareCreativeRequest) (*analyse.CompareCreative, error) {
	var resp analyse.CompareCreativeResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/product/analyse/compare_creative/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
