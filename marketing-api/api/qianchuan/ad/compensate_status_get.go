package ad

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/ad"
)

// CompensateStatusGet 获取计划成本保障状态
func CompensateStatusGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *ad.CompensateStatusGetRequest) ([]ad.CompensateStatus, error) {
	var resp ad.CompensateStatusGetResponse
	err := clt.Get(ctx, "v1.0/qianchuan/ad/compensate_status/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
