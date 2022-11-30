package v3

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	ad "github.com/bububa/oceanengine/marketing-api/model/ad/v3"
)

// Delete 批量删除广告
func Delete(clt *core.SDKClient, accessToken string, req *ad.DeleteRequest) (*ad.UpdateResponseData, error) {
	var resp ad.UpdateResponse
	err := clt.Post("v3.0/promotion/delete/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
