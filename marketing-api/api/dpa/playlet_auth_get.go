package dpa

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dpa"
)

// PlayletAuthGet 查询短剧商品原片授权申请状态
func PlayletAuthGet(clt *core.SDKClient, accessToken string, req *dpa.PlayletAuthGetRequest) (*dpa.PlayletAuthGetResult, error) {
	var resp dpa.PlayletAuthGetResponse
	if err := clt.GetAPI("2/dpa/playlet/auth/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
