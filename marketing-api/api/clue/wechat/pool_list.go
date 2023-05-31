package wechat

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/clue/wechat"
)

// PoolList 获取微信库微信号列表
func PoolList(clt *core.SDKClient, accessToken string, req *wechat.PoolListRequest) (*wechat.PoolListData, error) {
	var resp wechat.PoolListResponse
	if err := clt.Get("2/clue/wechat_pool/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
