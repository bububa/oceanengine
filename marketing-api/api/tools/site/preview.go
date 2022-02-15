package site

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/site"
)

// Preview 获取橙子建站站点预览地址
// 通过此接口，用户可以获取已创建站点的预览地址
// 预览地址有效期：20分钟
func Preview(clt *core.SDKClient, accessToken string, req *site.PreviewRequest) (string, error) {
	var resp site.PreviewResponse
	if err := clt.Get("2/tools/site/preview/", req, &resp, accessToken); err != nil {
		return "", err
	}
	return resp.Data.URL, nil
}
