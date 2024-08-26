package eventmanager

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/eventmanager"
)

// ShareCancel 事件管理资产取消共享
func ShareCancel(ctx context.Context, clt *core.SDKClient, accessToken string, req *eventmanager.ShareRequest) ([]eventmanager.ShareError, error) {
	var resp eventmanager.ShareResponse
	if err := clt.PostAPI(ctx, "v3.0/event_manager/share/cancel/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.ErrorList, nil
}
