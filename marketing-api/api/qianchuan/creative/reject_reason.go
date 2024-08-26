package creative

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/creative"
)

// RejectReason 获取创意审核建议
// 通过此接口用于获取广告创意纬度的审核意见
func RejectReason(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.RejectReasonRequest) ([]creative.RejectReasonList, error) {
	var resp creative.RejectReasonResponse
	err := clt.Get(ctx, "v1.0/qianchuan/creative/reject_reason/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
