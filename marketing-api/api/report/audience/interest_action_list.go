package audience

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/report/audience"
)

// InterestActionList 行为兴趣数据
// 此接口用于获取：行为+兴趣数据，如果你在计划的受众定向中设置了行为/兴趣定向，那么在这能看到所选的行为/兴趣指标所对应的数据，包括花费、展示、点击等。
// 行为兴趣数据不支持获取当天的数据；
func InterestActionList(ctx context.Context, clt *core.SDKClient, accessToken string, req *audience.ListRequest) (*audience.ListResponseData, error) {
	var resp audience.ListResponse
	err := clt.Get(ctx, "2/report/audience/interest_action/list/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
