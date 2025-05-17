package conversion

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/conversion"
)

// Conversion 新版转化回传
func Conversion(ctx context.Context, clt *core.SDKClient, req *conversion.Request) (int, error) {
	var resp conversion.Response
	if err := clt.AnalyticsPost(ctx, "conversion", req, &resp); err != nil && resp.Code != 128 {
		return resp.Code, err
	}
	return resp.Code, nil
}
