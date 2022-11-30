package v3

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	v3 "github.com/bububa/oceanengine/marketing-api/model/ad/v3"
)

// Update 修改广告
func Update(clt *core.SDKClient, accessToken string, req *v3.UpdateRequest) (*v3.UpdateResponseData, error) {
	var resp v3.UpdateResponse
	if err := clt.Post("v3.0/promotion/update/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
