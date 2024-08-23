package eventmanager

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/eventmanager"
)

// Share 事件管理资产共享
func Share(clt *core.SDKClient, accessToken string, req *eventmanager.ShareRequest) ([]eventmanager.ShareError, error) {
	var resp eventmanager.ShareResponse
	if err := clt.PostAPI("v3.0/event_manager/share/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.ErrorList, nil
}
