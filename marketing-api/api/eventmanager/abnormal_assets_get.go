package eventmanager

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/eventmanager"
)

// AbnormalAssetsGet 获取异常应用资产列表
// 可通过接口查询确认广告主未完成融合归因方案改造的异常资产。
// 接入逻辑可参考巨量广告转化融合归因优化方案接入手册V2.0，接入中如有问题可通过【雅典娜工单】排查/联系对接销售，如您认可通过巨量平台完成归因也可接入转化SDK简化接入流程，详情参考「对外」巨量广告转化安卓端SDK使用说明 for 6.16.9手册
// 接口仅返回异常资产内容，数据更新间隔为T+1
func AbnormalAssetsGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *eventmanager.AbnormalAssetsGetRequest) (*eventmanager.AbnormalAssetsGetResult, error) {
	var resp eventmanager.AbnormalAssetsGetResponse
	if err := clt.GetAPI(ctx, "v3.0/event_manager/abnormal_assets/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
