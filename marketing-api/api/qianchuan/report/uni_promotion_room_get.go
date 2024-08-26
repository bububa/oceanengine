package report

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/report"
)

// UniPromotionRoomGet 获取全域推广直播间维度数据
func UniPromotionRoomGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.UniPromotionDimensionGetRequest) (*report.UniPromotionDimensionGetResult, error) {
	var resp report.UniPromotionDimensionGetResponse
	err := clt.Get(ctx, "v1.0/qianchuan/report/uni_promotion/dimension_data/room/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
