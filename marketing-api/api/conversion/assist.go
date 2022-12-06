package conversion

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/conversion"
)

// Assist 客户回传其巨量引擎助攻的转化
func Assist(clt *core.SDKClient, req *conversion.Request) error {
	var resp conversion.Response
	return clt.AnalyticsPost("assist", req, &resp)
}
