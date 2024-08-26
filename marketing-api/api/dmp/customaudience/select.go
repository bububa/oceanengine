package customaudience

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dmp/customaudience"
)

// Select 人群包列表
// 通过此接口你可以查询广告主下存在的的人群包列表和信息。
// 信息包括人群包的id，可用状态，来源，覆盖人群等； 该接口可以查到广告主id下所有人群包，无论是API创建的还是平台创建的。
// 当delivery_status:人群包可投放状态为可投放时，才可以在【创建广告计划】和【修改广告计划】的用户定向中使用！
func Select(ctx context.Context, clt *core.SDKClient, accessToken string, req *customaudience.SelectRequest) (*customaudience.SelectResponseData, error) {
	var resp customaudience.SelectResponse
	err := clt.Get(ctx, "2/dmp/custom_audience/select/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
