package audience

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/report/audience"
)

// Tag 兴趣数据
// 获取兴趣受众数据，如果你在计划的受众定向中设置了兴趣定向，那么在这能看到所选的兴趣指标所对应的数据，包括花费、展示、点击等。
// 受众数据不支持获取当天的数据；
func Tag(ctx context.Context, clt *core.SDKClient, accessToken string, req *audience.Request) ([]audience.ResponseData, error) {
	var resp audience.Response
	err := clt.Get(ctx, "2/report/audience/tag/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
