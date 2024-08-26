package v3

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/enum"
	v3 "github.com/bububa/oceanengine/marketing-api/model/eventmanager/v3"
)

// DeepBidTypeGet 获取可用深度优化方式体验版
func DeepBidTypeGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *v3.DeepBidTypeGetRequest) ([]enum.DeepBidType, error) {
	var resp v3.DeepBidTypeGetResponse
	if err := clt.GetAPI(ctx, "v3.0/event_manager/deep_bid_type/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.DeepBidType, nil
}
