package creative

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/creative"
)

// CustomCreativeCreate 创建自定义创意（营销链路）
// 创意包括了基础创意、创意素材、创意分类等信息，对于其中的概念解释可以参考：【广告创意】
// 创意创建涉及了多个资产管理：素材管理、落地页管理等；开发者需要提前根据开放接口构建这些资产功能，以免创建计划卡住

// 新版营销链路广告创意差异点

// 新版接口新增落地页，小程序等相关入参
// 下线推广卡片，附加创意等相关字段，替换为创意组件形式传入
// 投放位置与监测链接变更为计划接口传入
// 每个广告主每日最多创建500个创意（自定义创意+程序化创意）
// 其中视频的时长需要>=4s，否则会报错
// 对于不打算传的字段，不要传“”或者null，传了会校验
// 如果计划ID下已有创意信息，需要使用update_v2接口进行修改或者新增创意素材，否则会报错
func CustomCreativeCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.CustomCreativeCreateRequest) (*creative.CustomCreativeCreateResponseData, error) {
	var resp creative.CustomCreativeCreateResponse
	err := clt.Post(ctx, "2/creative/custom_creative/create/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
