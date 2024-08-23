package promotion

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/promotion"
)

// Create 创建广告
func Create(clt *core.SDKClient, accessToken string, req *promotion.CreateRequest) (uint64, error) {
	var resp promotion.CreateResponse
	if err := clt.PostAPI("v3.0/promotion/create/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.PromotionID, nil
}
