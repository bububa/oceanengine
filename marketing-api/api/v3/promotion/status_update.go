package promotion

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/promotion"
)

// StatusUpdate 更新广告状态
func StatusUpdate(clt *core.SDKClient, accessToken string, req *promotion.StatusUpdateRequest) (*promotion.UpdateResponseData, error) {
	var resp promotion.UpdateResponse
	err := clt.Post("v3.0/promotion/status/update/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
