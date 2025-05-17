package report

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/report"
)

// MistyGet 分级模糊数据
// 此接口用于第三方服务商获取广告主相关的数据信息，这些指标数据以分级区间的形式展示。
// 指标数据分级区间值示例如下：[0 ,0]，(1, 2]，(2, +)。[ ]表示闭区间；( )表示开区间；+表示正无穷。通常返回左开右闭区间的数据。
// 目前分级模糊数据接口只支持获取基础报表、受众分析报表、关键词报表、搜索词报表。
func MistyGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.IntegratedRequest) (*report.IntegratedResponseData, error) {
	var resp report.IntegratedResponse
	err := clt.Get(ctx, "2/report/misty/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
