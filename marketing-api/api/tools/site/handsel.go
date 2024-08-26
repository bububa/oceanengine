package site

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/site"
)

// Handsel 建站工具-建站转赠
// 通过此接口，用户可以实现站点的转赠功能，将某一广告主的站点共享给其他特定的广告主，转赠成功后，广告主们拥有站点的共同使用权。 不限制主体，同广告主不能进行转赠操作 注意
// 用户在传入请求参数site_ids站点id列表时，请确保id的正确性，存在错误将导致转赠全部失败！
func Handsel(ctx context.Context, clt *core.SDKClient, accessToken string, req *site.HandselRequest) (*site.HandselResponseData, error) {
	var resp site.HandselResponse
	if err := clt.Post(ctx, "2/tools/site/handsel/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
