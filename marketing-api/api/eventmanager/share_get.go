package eventmanager

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/eventmanager"
)

// ShareGet 事件管理资产查看共享范围
func ShareGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *eventmanager.ShareGetRequest) (*eventmanager.ShareGetData, error) {
	var resp eventmanager.ShareGetResponse
	err := clt.GetAPI(ctx, "v3.0/event_manager/share/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
