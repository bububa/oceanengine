package v3

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	v3 "github.com/bububa/oceanengine/marketing-api/model/tools/adraise/v3"
)

// StatusGet 获取一键起量方案列表
func StatusGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *v3.StatusGetRequest) ([]v3.PromotionRaiseStatus, error) {
	var resp v3.StatusGetResponse
	if err := clt.GetAPI(ctx, "v3.0/tools/promotion_raise_status/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
