package subscribe

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/subscribe"
)

// AccountsList 查询订阅 Adv
// 查询订阅任务所订阅的 Adv
func AccountsList(ctx context.Context, clt *core.SDKClient, accessToken string, req *subscribe.AccountsListRequest) (*subscribe.AccountsListResult, error) {
	var resp subscribe.AccountsListResponse
	if err := clt.GetAPI(ctx, "v3.0/subscribe/accounts/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
