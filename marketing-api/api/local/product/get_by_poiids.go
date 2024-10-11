package product

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/local/product"
)

// GetByPoiIDs 根据门店ID查询门店下商品ID
// 该接口使用场景说明：
// 通过【获取抖音主页视频】接口拉取挂载了商品锚点的视频时，需要传入商品ids。如果为推门店场景，可以通过该接口获取所推广门店下的商品ID，进而拉取到挂载了门店下商品锚点的抖音主页视频进行投放。
// 仅创编推广目的=团购成交/门店引流/内容加热的计划时，可能涉及该接口的使用；推广目的=获取线索的计划不涉及。
func GetByPoiIDs(ctx context.Context, clt *core.SDKClient, accessToken string, req *product.GetByPoiIDsRequest) ([]uint64, error) {
	var resp product.GetByPoiIDsResponse
	if err := clt.GetAPI(ctx, "v3.0/local/product/get_by_poiids/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.ProductIDs, nil
}
