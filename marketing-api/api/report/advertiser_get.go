package report

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/report"
)

// AdvertiserGet 广告主数据
// 此接口用于获取广告账户纬度的投放数据，包括消耗、点击、展示等指标，具体可以参考应答参数指标说明。
// 广告主消耗的变化来判断组、计划、创意消耗是否发生了变化，再定时拉取细纬度数据以免做无效的数据调用，减少自身服务的压力；
func AdvertiserGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.GetRequest) (*report.GetResponseData, error) {
	var resp report.GetResponse
	err := clt.Get(ctx, "2/report/advertiser/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
