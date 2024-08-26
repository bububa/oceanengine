package ad

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/ad"
)

// RejectReason 获取计划审核建议
func RejectReason(ctx context.Context, clt *core.SDKClient, accessToken string, req *ad.RejectReasonRequest) ([]ad.RejectReason, error) {
	var resp ad.RejectReasonResponse
	if err := clt.Get(ctx, "2/ad/reject_reason/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
