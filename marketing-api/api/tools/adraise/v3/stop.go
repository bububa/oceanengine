package v3

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	v3 "github.com/bububa/oceanengine/marketing-api/model/tools/adraise/v3"
)

// Stop 关停一键起量
func Stop(ctx context.Context, clt *core.SDKClient, accessToken string, req *v3.StopRequest) error {
	return clt.PostAPI(ctx, "v3.0/tools/promotion_raise/stop/", req, nil, accessToken)
}
