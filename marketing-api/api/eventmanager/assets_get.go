package eventmanager

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/eventmanager"
)

// AssetsGet 获取已创建资产列表
func AssetsGet(clt *core.SDKClient, accessToken string, req *eventmanager.AssetsGetRequest) (*eventmanager.AssetsGetResponseData, error) {
	var resp eventmanager.AssetsGetResponse
	err := clt.Get("2/tools/event/assets/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
