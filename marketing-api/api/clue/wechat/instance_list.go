package wechat

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/clue/wechat"
)

// InstanceList 获取微信号码包列表
func InstanceList(clt *core.SDKClient, accessToken string, req *wechat.InstanceListRequest) (*wechat.InstanceListData, error) {
	var resp wechat.InstanceListResponse
	if err := clt.Get("2/clue/wechat_instance/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
