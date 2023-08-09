package enterprise

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/enterprise"
)

// ListItem 获取企业号视频列表
func ListItem(clt *core.SDKClient, accessToken string, req *enterprise.ListItemRequest) (*enterprise.ListItemResponse, error) {
	var resp enterprise.ListItemResponse
	err := clt.Get("v1.0/enterprise/item/list/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
