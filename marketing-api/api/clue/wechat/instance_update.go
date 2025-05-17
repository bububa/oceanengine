package wechat

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/clue/wechat"
)

// InstanceUpdate 更新微信号码包
func InstanceUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *wechat.InstanceUpdateRequest) (*wechat.InstanceUpdateResult, error) {
	var resp wechat.InstanceUpdateResponse
	if err := clt.Post(ctx, "2/clue/wechat_instance/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
