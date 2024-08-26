package sitetemplate

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/sitetemplate"
)

// PicURLGet 获取模板/站点URL
// 通过 site_id / template_id 获取站点/模板下的图片加签URL
func PicURLGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *sitetemplate.PicURLGetRequest) (map[string]string, error) {
	var resp sitetemplate.PicURLGetResponse
	if err := clt.Get(ctx, "2/tools/site_template/pic_url/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.URLMap, nil
}
