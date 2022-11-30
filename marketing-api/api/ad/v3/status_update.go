package v3

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	ad "github.com/bububa/oceanengine/marketing-api/model/ad/v3"
)

// StatusUpdate 更新计划状态
func StatusUpdate(clt *core.SDKClient, accessToken string, req *ad.StatusUpdateRequest) (*ad.UpdateResponseData, error) {
	var resp ad.UpdateResponse
	err := clt.Post("v3.0/promotion/status/update/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
