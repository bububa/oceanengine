package nativeanchor

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/nativeanchor"
)

// Create 原生锚点创建
func Create(clt *core.SDKClient, accessToken string, req *nativeanchor.CreateRequest) (*nativeanchor.CreateResponseData, error) {
	var resp nativeanchor.CreateResponse
	if err := clt.Post("v3.0/tools/native_anchor/create/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
