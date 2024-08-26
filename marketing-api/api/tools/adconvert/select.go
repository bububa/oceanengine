package adconvert

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/adconvert"
)

// Select 转化目标列表
// 查询广告主名义下选定的转化目标信息列表
// 可通过convert_ids查询指定ID的转化目标，或通过opt_status查询指定状态的转化目标
// 如果是需要查询广告计划可用的转化请使用接口【查询计划可用转化目标】
// 转化convert_ids为选填，如果不传，应答参数中默认返回广告主下所有的转化目标ID
func Select(ctx context.Context, clt *core.SDKClient, accessToken string, req *adconvert.SelectRequest) (*adconvert.SelectResponseData, error) {
	var resp adconvert.SelectResponse
	err := clt.Get(ctx, "2/tools/ad_convert/select/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
