package creative

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/creative"
)

// RejectReason 获取创意审核建议
// 此接口用于获取创意审核建议；
// 所有的程序化创意都是审核通过的；
// 如果获取创意遇到问题，可通过 常见问题 来解决
func RejectReason(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.RejectReasonRequest) ([]creative.RejectReason, error) {
	var resp creative.RejectReasonResponse
	if err := clt.Get(ctx, "2/creative/reject_reason/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
