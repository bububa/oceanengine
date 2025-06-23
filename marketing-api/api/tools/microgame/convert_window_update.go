package microgame

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/microgame"
)

// ConvertWindowUpdate 修改字节小游戏归因激活时间窗
// 使用场景：当前接口仅支持修改字节小游戏归因激活时间窗，表示已激活用户最后一次打开小程序后，时隔多久可被重新判定为未激活用户，该周期针对字节自归因的广告转化目标生效，默认为30天，可通过该接口编辑修改。
//
// 请注意：
// 每个小游戏1自然日内只能设置1次去重周期；
// 该接口使用前需要找对应销售/运营同学加白才可使用功能；
// 目前接口功能灰度中，若报错显示无权限，可联系运营/销售操作功能开放；
func ConvertWindowUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *microgame.ConvertWindowUpdateRequest) error {
	return clt.PostAPI(ctx, "v3.0/tools/micro_game/convert_window/update/", req, nil, accessToken)
}
