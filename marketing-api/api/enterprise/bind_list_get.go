package enterprise

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/enterprise"
)

// BindListGet 获取广告主关联的企业号列表
func BindListGet(clt *core.SDKClient, accessToken string, req *enterprise.BindListGetRequest) ([]enterprise.BindItem, error) {
	var resp enterprise.BindListGetResponse
	if err := clt.Get("v1.0/enterprise/bind/list/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
