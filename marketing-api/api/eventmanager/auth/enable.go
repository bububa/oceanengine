package auth

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/eventmanager/auth"
)

// Enable 开启鉴权
// 广告主开启鉴权。
// 鉴权关闭状态下，鉴权失败的转化数据不会被丢弃，同时返回鉴权失败原因，方便广告主调试定位。开启鉴权后，鉴权失败的转化数据将被丢弃。平台会缓存广告主的鉴权开启状态，缓存时长约3分钟，这意味着，最长会有3分钟的过渡时间。
// 广告主确认上线后有成功回传转化，且没有鉴权失败（观察24h以上，最长不超过一周)之后，可请求“开启鉴权”接口，平台检查是否满足开启条件，如果满足则开启，否则开启失败。
func Enable(ctx context.Context, clt *core.SDKClient, accessToken string, advertiserID uint64) error {
	req := auth.EnableRequest{
		AdvertiserID: advertiserID,
	}
	return clt.Post(ctx, "2/event_manager/auth/enable", &req, nil, accessToken)
}
