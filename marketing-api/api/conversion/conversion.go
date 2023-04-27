package conversion

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/conversion"
)

// Conversion 新版转化回传
func Conversion(clt *core.SDKClient, req *conversion.Request) (int, error) {
	var resp conversion.Response
	if err := clt.AnalyticsPost("conversion", req, &resp); err != nil && resp.Code != 128 {
		return resp.Code, err
	}
	return resp.Code, nil
}
