package star

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/star"
)

// DemandList 获取星图客户任务列表
// 此接口用于根据星图ID，获取星图客户账号下的全部任务，包含星图任务id、任务名称、结算方式、组件类型、任务创建时间。
func DemandList(ctx context.Context, clt *core.SDKClient, accessToken string, req *star.DemandListRequest) (*star.DemandListResult, error) {
	var resp star.DemandListResponse
	err := clt.Get(ctx, "2/star/demand/list/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
