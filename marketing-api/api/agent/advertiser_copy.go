package agent

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/agent"
)

// AdvertiserCopy 广告主账户复制
func AdvertiserCopy(clt *core.SDKClient, accessToken string, req *agent.AdvertiserCopyRequest) (*agent.AdvertiserCopyResult, error) {
	var resp agent.AdvertiserCopyResponse
	if err := clt.Post("2/agent/advertiser/copy/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
