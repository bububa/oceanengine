package v3

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/adraise/v3"
)

// StatusGet 获取一键起量方案列表
func StatusGet(clt *core.SDKClient, accessToken string, req *v3.StatusGetRequest) ([]v3.PromotionRaiseStatus, error) {
	var resp v3.StatusGetResponse
	if err := clt.Get("v3.0/tools/promotion_raise_status/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
