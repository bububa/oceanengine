package agent

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/agent"
)

// QueryBookingBusinessEntityIDGet 排期—查询业务实体ID
func QueryBookingBusinessEntityIDGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *agent.QueryBookingBusinessEntityIDGetRequest) ([]agent.BusinessEntityIDInfo, error) {
	var resp agent.QueryBookingBusinessEntityIDGetResponse
	if err := clt.GetAPI(ctx, "2/query/booking/business_entity_id/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.BusinessEntityIDInfos, nil
}
