package v3

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	v3 "github.com/bububa/oceanengine/marketing-api/model/report/v3"
)

// MaterialGet 素材数据报表
func MaterialGet(clt *core.SDKClient, accessToken string, req *v3.MaterialGetRequest) (*v3.MaterialGetResult, error) {
	var resp v3.MaterialGetResponse
	if err := clt.Get("v3.0/report/material/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
