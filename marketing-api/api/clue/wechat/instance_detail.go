package wechat

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/clue/wechat"
)

// InstanceDetail 获取微信号码包详情
func InstanceDetail(clt *core.SDKClient, accessToken string, req *wechat.InstanceDetailRequest) (*wechat.Instance, error) {
	var resp wechat.InstanceDetailResponse
	if err := clt.Get("2/clue/wechat_instance/detail/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
