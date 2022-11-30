package v3

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	v3 "github.com/bububa/oceanengine/marketing-api/model/ad/v3"
)

// List 获取广告列表
func List(clt *core.SDKClient, accessToken string, req *v3.ListRequest) (*v3.ListResponseData, error) {
	var resp v3.ListResponse
	if err := clt.Get("v3.0/promotion/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
