package nativeanchor

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/nativeanchor"
)

// Delete 删除原生锚点
func Delete(ctx context.Context, clt *core.SDKClient, accessToken string, req *nativeanchor.DeleteRequest) error {
	return clt.PostAPI(ctx, "v3.0/native_anchor/delete/", req, nil, accessToken)
}
