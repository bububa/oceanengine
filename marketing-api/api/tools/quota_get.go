package tools

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools"
)

// QuotaGet 查询在投计划配额
// 该接口用于查询广告账户的在投计划配额和使用进度
func QuotaGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *tools.QuotaGetRequest) (*tools.QuotaGetResult, error) {
	var resp tools.QuotaGetResponse
	if err := clt.Get(ctx, "2/tools/quota/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
