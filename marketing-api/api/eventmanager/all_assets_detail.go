package eventmanager

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/eventmanager"
)

// AllAssetsDetail 获取已创建资产详情（新）
// - 本接口为巨量引擎开放平台2024年6月5日新增接口，接口能力覆盖旧接口「获取已创建资产列表」，且新增支持以下能力：
//   - 支持获取橙子落地页TETRIS_EXTERNAL、应用APP类型的资产详情信息
//   - 可通过资产ID筛选查询资产详情
//
// - 「获取已创建资产列表」接口即将于8月6日下线，可切换至本接口获取资产详情
// - 应答参数返回的资产范围如下：
//   - 仅支持查询账户下创建以及共享中状态、未删除的资产详情，共享失败的资产不支持查询
//   - 账户下不返回已删除资产的资产详情信息
func AllAssetsDetail(ctx context.Context, clt *core.SDKClient, accessToken string, req *eventmanager.AllAssetsDetailRequest) ([]eventmanager.AssetDetail, error) {
	var resp eventmanager.AllAssetsDetailResponse
	if err := clt.GetAPI(ctx, "2/tools/event/all_assets/detail/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.AssetList, nil
}
