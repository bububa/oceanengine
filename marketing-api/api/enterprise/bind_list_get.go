package enterprise

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/enterprise"
)

// BindListGet 获取广告主关联的企业号列表
func BindListGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *enterprise.BindListGetRequest) ([]enterprise.BindItem, error) {
	var resp enterprise.BindListGetResponse
	if err := clt.Get(ctx, "v1.0/enterprise/bind/list/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
