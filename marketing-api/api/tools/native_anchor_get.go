package tools

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools"
)

// NativeAnchorGet 获取账户下原生锚点
func NativeAnchorGet(clt *core.SDKClient, accessToken string, req *tools.NativeAnchorGetRequest) (*tools.NativeAnchorGetResponseData, error) {
	var resp tools.NativeAnchorGetResponse
	if err := clt.Get("v3.0/tools/native_anchor/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
