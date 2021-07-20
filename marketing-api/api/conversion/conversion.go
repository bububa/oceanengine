package conversion

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/conversion"
)

// 新版转化回传
func Conversion(clt *core.SDKClient, req *conversion.Request) error {
	var resp conversion.Response
	return clt.AnalyticsPost("conversion", req, &resp)
}
