package promotion

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/promotion"
)

// RejectReason 获取计划审核建议
func RejectReason(ctx context.Context, clt *core.SDKClient, accessToken string, req *promotion.RejectReasonRequest) ([]promotion.RejectReason, error) {
	var resp promotion.RejectReasonResponse
	if err := clt.GetAPI(ctx, "v3.0/promotion/reject_reason/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
