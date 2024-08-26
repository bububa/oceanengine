package tools

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools"
)

// GrayGet 查询白名单能力
// 支持客户通过接口查询广告主ID是否命中各项白名单功能
func GrayGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *tools.GrayGetRequest) ([]tools.GrayItem, error) {
	var resp tools.GrayGetResponse
	if err := clt.GetAPI(ctx, "v3.0/tools/gray/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
