package unipromotion

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	unipromotion "github.com/bububa/oceanengine/marketing-api/model/qianchuan/uni_promotion"
)

// StatusUpdate 更改全域推广计划状态
func StatusUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *unipromotion.StatusUpdateRequest) ([]unipromotion.StatusUpdateResult, error) {
	var resp unipromotion.StatusUpdateResponse
	if err := clt.Post(ctx, "v1.0/qianchuan/uni_promotion/ad/status/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.Results, nil
}
