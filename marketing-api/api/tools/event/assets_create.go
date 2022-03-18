package event

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/event"
)

// AssetsCreate 创建资产
// 此接口用于创建广告账户下资产
func AssetsCreate(clt *core.SDKClient, accessToken string, req *event.AssetsCreateRequest) (uint64, error) {
	var resp event.AssetsCreateResponse
	if err := clt.Post("2/tools/event_manager/assets/create/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.AssetID, nil
}
