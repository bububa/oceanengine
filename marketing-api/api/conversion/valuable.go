package conversion

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/conversion"
)

// Valuable 客户回传其他渠道全部转化事件，由巨量引擎判定是否属于可应用的助攻事件
func Valuable(ctx context.Context, clt *core.SDKClient, req *conversion.Request) error {
	var resp conversion.Response
	return clt.AnalyticsPost(ctx, "valuable", req, &resp)
}
