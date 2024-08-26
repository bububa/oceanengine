package site

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/site"
)

// Copy 建站工具-建站复制
// 通过此接口，用户可以实现站点的复制功能，成功后生成一个新站点id，站点内容和原站点一致。 注意
// 用户在传入请求参数site_ids站点id列表时，需确保id的正确性，存在错误将导致复制全部失败！
func Copy(ctx context.Context, clt *core.SDKClient, accessToken string, req *site.CopyRequest) (*site.CopyResponseData, error) {
	var resp site.CopyResponse
	if err := clt.Post(ctx, "2/tools/site/copy/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
