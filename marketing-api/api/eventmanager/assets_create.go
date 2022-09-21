package eventmanager

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/eventmanager"
)

// AssetsCreate 创建资产
// 此接口用于创建广告账户下资产
func AssetsCreate(clt *core.SDKClient, accessToken string, req *eventmanager.AssetsCreateRequest) (uint64, error) {
	var resp eventmanager.AssetsCreateResponse
	if err := clt.Post("2/event_manager/assets/create/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.AssetID, nil
}
