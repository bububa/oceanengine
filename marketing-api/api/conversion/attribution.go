package conversion

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/conversion"
)

// Attribution 电话转化回传
func Attribution(ctx context.Context, clt *core.SDKClient, req *conversion.Request) error {
	var resp conversion.Response
	return clt.AnalyticsV1Post(ctx, "attribution", req, &resp)
}
