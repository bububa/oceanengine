package nativeanchor

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/nativeanchor"
)

// Get 获取账户下原生锚点
func Get(clt *core.SDKClient, accessToken string, req *nativeanchor.GetRequest) (*nativeanchor.GetResponseData, error) {
	var resp nativeanchor.GetResponse
	if err := clt.Get("v3.0/tools/native_anchor/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
