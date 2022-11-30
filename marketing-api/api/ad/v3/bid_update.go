package v3

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	ad "github.com/bububa/oceanengine/marketing-api/model/ad/v3"
)

// BidUpdate 更新出价
func BidUpdate(clt *core.SDKClient, accessToken string, req *ad.BidUpdateRequest) (*ad.UpdateResponseData, error) {
	var resp ad.UpdateResponse
	err := clt.Post("v3.0/promotion/bid/update/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
