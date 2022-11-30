package v3

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	v3 "github.com/bububa/oceanengine/marketing-api/model/ad/v3"
)

// Create 创建广告
func Create(clt *core.SDKClient, accessToken string, req *v3.CreateRequest) (uint64, error) {
	var resp v3.CreateResponse
	if err := clt.Post("v3.0/promotion/create/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.PromotionID, nil
}
