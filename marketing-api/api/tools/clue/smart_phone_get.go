package clue

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/clue"
)

// SmartPhoneGet 建站工具——查询已有智能电话
// 通过此接口，用户可以获取已有智能电话实例列表，可用于建站。
// 列表栏包括智能电话实例ID、实例名称、是否为真实可接通的电话、电话id、电话号码、创建时间等信息。
func SmartPhoneGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *clue.SmartPhoneGetRequest) (*clue.SmartPhoneGetResponseData, error) {
	var resp clue.SmartPhoneGetResponse
	if err := clt.Get(ctx, "2/tools/clue/smart_phone/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
