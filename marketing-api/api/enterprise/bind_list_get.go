package enterprise

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/enterprise"
)

// GetBindList 获取广告主关联的企业号列表
func GetBindList(clt *core.SDKClient, accessToken string, req *enterprise.GetBindListRequest) ([]*enterprise.BindItem, error) {
	var resp enterprise.GetBindListResponse
	err := clt.Get("v1.0/enterprise/bind/list/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
