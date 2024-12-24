package dpa

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dpa"
)

// ClueProductDelete 删除升级版商品
// 【使用场景】对广告主创建的「升级版」商品进行删除操作，注意已经关联计划的商品不允许进行删除，支持批量，一次性调用最大个数为100，服务为部分成功部分失败。
func ClueProductDelete(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.ClueProductDeleteRequest) (*dpa.ClueProductDeleteResult, error) {
	var resp dpa.ClueProductDeleteResponse
	if err := clt.PostAPI(ctx, "2/dpa/clue_product/delete/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
