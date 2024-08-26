package auth

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/eventmanager/auth"
)

// Disable 关闭鉴权
// 广告主关闭鉴权。关闭鉴权后，回传的转化数据即使鉴权失败也不会被丢弃。平台会缓存广告主的鉴权开启状态，缓存时长约3分钟，这意味着，最长会有3分钟的过渡时间。
func Disable(ctx context.Context, clt *core.SDKClient, accessToken string, advertiserID uint64) error {
	req := auth.DisableRequest{
		AdvertiserID: advertiserID,
	}
	return clt.Post(ctx, "2/event_manager/auth/disable", &req, nil, accessToken)
}
