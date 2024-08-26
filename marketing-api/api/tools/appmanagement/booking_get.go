package appmanagement

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/appmanagement"
)

// BookingGet 查询游戏信息
// 查询游戏预约列表接口
func BookingGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *appmanagement.AppListRequest) (*appmanagement.AppListResponseData, error) {
	var resp appmanagement.AppListResponse
	if err := clt.Get(ctx, "2/tools/app_management/booking/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
