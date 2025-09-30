package unipromotion

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	unipromotion "github.com/bububa/oceanengine/marketing-api/model/qianchuan/uni_promotion"
)

// AdBudgetUpdate 更新全域推广计划预算
func AdBudgetUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *unipromotion.AdBudgetUpdateRequest) ([]unipromotion.AdUpdateResult, error) {
	var resp unipromotion.AdUpdateResponse
	if err := clt.Post(ctx, "v1.0/qianchuan/uni_promotion/ad/name/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.Results, nil
}
