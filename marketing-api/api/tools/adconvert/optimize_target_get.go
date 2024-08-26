package adconvert

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/adconvert"
)

// OptimizeTargetGet 查询广告计划可用优化目标
// 查询ocpc、ocpm广告计划可使用的转化目标ID以及对应的转化目标信息（转化目标的相关属性，不包括转化跟踪效果数据） 该接口支持的推广目的有：应用推广、销售线索收集、快应用、抖音号、直播间
// 影响因素为推广目的、投放范围、投放内容以及附加创意 查询结果都是可用状态的转化，均为激活状态的转化
// 「convert_id」仅返回自定义转化目标ID，预定义转化目标返回为Null，对应数字值可根据external_action参考【枚举值-转化类型】
// 游戏预约场景目前只支持落地页，游戏预约场景必传落地页链接external_url
// 快应用不支持IOS链接
// 下载方式为下载链接或者快应用+下载链接时，app_type必填
// app_type为APP_ANDROID，package_name必填
// 下载方式为落地页时，external_url必填
func OptimizeTargetGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *adconvert.OptimizeTargetGetRequest) ([]adconvert.OptimizeTarget, error) {
	var resp adconvert.OptimizeTargetGetResponse
	err := clt.Get(ctx, "2/tools/ad_convert/optimize_target/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}
