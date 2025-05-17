package v3

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	v3 "github.com/bububa/oceanengine/marketing-api/model/report/v3"
)

// PromotionGet 广告数据报表
func PromotionGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *v3.PromotionGetRequest) (*v3.PromotionGetResult, error) {
	var resp v3.PromotionGetResponse
	if err := clt.GetAPI(ctx, "v3.0/report/promotion/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
