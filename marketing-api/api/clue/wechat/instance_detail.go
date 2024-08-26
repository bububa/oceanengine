package wechat

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/clue/wechat"
)

// InstanceDetail 获取微信号码包详情
func InstanceDetail(ctx context.Context, clt *core.SDKClient, accessToken string, req *wechat.InstanceDetailRequest) (*wechat.Instance, error) {
	var resp wechat.InstanceDetailResponse
	if err := clt.Get(ctx, "2/clue/wechat_instance/detail/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
