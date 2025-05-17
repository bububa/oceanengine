package report

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/report"
)

// CreativeGet 广告创意数据
// 此接口用于获取广告创意纬度的投放数据，包括消耗、点击、展示等指标，具体可以参考应答参数指标说明。
// 默认不获取删除广告组数据，如果需要查询删除的数据，可以在filtering中设置status，传AD_STATUS_ALL（所有包含已删除）
// 如果一次查询的广告创意个数超过5000个，会被截断，具体说明参考，文档下方-常见问题
// 需要分批获取自定义创意与非自定义创意，一次只能取一种类型创意数据，具体可以参考creative_material_modes字段说明；
func CreativeGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.GetRequest) (*report.GetResponseData, error) {
	var resp report.GetResponse
	err := clt.Get(ctx, "2/report/creative/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
