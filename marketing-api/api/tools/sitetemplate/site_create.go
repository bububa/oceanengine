package sitetemplate

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/sitetemplate"
)

// SiteCreate 基于模板创建站点
// 可以基于已创建的模版新建或者编辑落地页站点，如需进行发布请调用【更改橙子建站站点状态】接口
// 未发布站点无法使用
func SiteCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *sitetemplate.SiteCreateRequest) (uint64, error) {
	var resp sitetemplate.SiteCreateResponse
	if err := clt.Post(ctx, "2/tools/site_template/site/create/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.SiteID, nil
}
