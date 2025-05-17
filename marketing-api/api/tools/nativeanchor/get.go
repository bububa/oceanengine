package nativeanchor

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/nativeanchor"
)

// Get 获取账户下原生锚点
func Get(ctx context.Context, clt *core.SDKClient, accessToken string, req *nativeanchor.GetRequest) (*nativeanchor.GetResponseData, error) {
	var resp nativeanchor.GetResponse
	if err := clt.GetAPI(ctx, "v3.0/native_anchor/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
