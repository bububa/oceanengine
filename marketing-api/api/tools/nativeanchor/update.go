package nativeanchor

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/nativeanchor"
)

// Update 更新原生锚点
func Update(clt *core.SDKClient, accessToken string, req *nativeanchor.UpdateRequest) (*nativeanchor.UpdateResponseData, error) {
	var resp nativeanchor.UpdateResponse
	if err := clt.Post("v3.0/native_anchor/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
