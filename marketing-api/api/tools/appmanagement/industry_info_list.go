package appmanagement

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/appmanagement"
)

// IndustryInfoList 获取应用细分分类及题材标签
// 获取应用细分分类及题材标签「支持所有账户体系」，查询所有应用细分分类及题材标签情况。
func IndustryInfoList(ctx context.Context, clt *core.SDKClient, accessToken string, req *appmanagement.IndustryInfoListRequest) ([]appmanagement.Industry, error) {
	var resp appmanagement.IndustryInfoListResponse
	if err := clt.Get(ctx, "2/tools/app_management/industry_info/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.Industries, nil
}
