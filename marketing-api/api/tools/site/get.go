package site

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/site"
)

// Get 获取橙子建站站点列表
// 通过此接口，用户可以获取广告主建站列表。
// 列表栏包括建站ID、建站名称、建站状态、建站类型、建站类别、站点缩略图等信息。
// 该接口当前还不会返回建站地址，建站地址可由如下格式拼装得到：https://www.chengzijianzhan.com/tetris/page/XXXXXXXXXXXXX/（其中XX是建站ID，拼装后就可获得投放的URL/预览URL）
// 鲁班电商广告主在投放以“商品推广”推广为目的广告时，如使用橙子建站落地页，只能使用“多品页”类型。同时，需要确保"status"站点状态为“auditAccepted”审核通过。
func Get(ctx context.Context, clt *core.SDKClient, accessToken string, req *site.GetRequest) (*site.GetResponseData, error) {
	var resp site.GetResponse
	if err := clt.Get(ctx, "2/tools/site/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
