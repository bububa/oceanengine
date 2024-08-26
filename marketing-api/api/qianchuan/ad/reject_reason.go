package ad

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/ad"
)

// RejectReason 获取计划审核建议
// 通过此接口用于获取广告计划纬度的审核意见
func RejectReason(ctx context.Context, clt *core.SDKClient, accessToken string, req *ad.RejectReasonRequest) ([]ad.RejectReasonList, error) {
	var resp ad.RejectReasonResponse
	err := clt.Get(ctx, "v1.0/qianchuan/ad/reject_reason/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
