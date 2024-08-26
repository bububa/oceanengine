package agent

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/agent"
)

// AdvertiserUpdate 修改广告主
// 修改广告主信息，可更改内容包括账户名称、联系人、手机号码、固定电话，除此之外其他内容不允许修改。
func AdvertiserUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *agent.AdvertiserUpdateRequest) (*agent.AdvertiserUpdateResponseData, error) {
	var resp agent.AdvertiserUpdateResponse
	if err := clt.Post(ctx, "2/agent/advertiser/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
