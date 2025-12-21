package unipromotion

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	unipromotion "github.com/bububa/oceanengine/marketing-api/model/qianchuan/uni_promotion"
)

// AdScheduleDateUpdate 更新全域推广计划投放时间
func AdScheduleDateUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *unipromotion.AdScheduleDateUpdateRequest) ([]unipromotion.AdUpdateResult, error) {
	var resp unipromotion.AdUpdateResponse
	if err := clt.Post(ctx, "v1.0/qianchuan/uni_promotion/ad/schedule_date/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.Results, nil
}
