package report

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/report"
)

// AdvertiserGet 广告主数据
// 当前广告账户数据仅包含“通投广告”部分数据，暂不包含搜索广告数据
func AdvertiserGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.GetRequest) (*report.GetResponseData, error) {
	var resp report.GetResponse
	err := clt.Get(ctx, "v1.0/qianchuan/report/advertiser/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
