package advertiser

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/advertiser"
)

// TypeGet 获取千川账户类型
func TypeGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *advertiser.TypeGetRequest) ([]advertiser.Advertiser, error) {
	var resp advertiser.TypeGetResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/advertiser/type/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
