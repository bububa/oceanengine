package creative

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/creative"
)

// Get 获取创意列表
func Get(clt *core.SDKClient, accessToken string, req *creative.GetRequest) (*creative.GetResponseData, error) {
	var resp creative.GetResponse
	err := clt.Get("2/creative/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
