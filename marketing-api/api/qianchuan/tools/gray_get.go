package tools

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	qianchuanTools "github.com/bububa/oceanengine/marketing-api/model/qianchuan/tools"
	"github.com/bububa/oceanengine/marketing-api/model/tools"
)

// GrayGet 查询白名单能力
// 支持客户通过接口查询广告主ID是否命中各项白名单功能
func GrayGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *tools.GrayGetRequest) (*qianchuanTools.GrayGetResult, error) {
	var resp qianchuanTools.GrayGetResponse
	if err := clt.GetAPI(ctx, "v3.0/qianchuan/tools/gray/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
