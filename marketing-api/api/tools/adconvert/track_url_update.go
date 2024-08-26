package adconvert

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/adconvert"
)

// TrackURLUpdateStatus 修改转化监测链接
// 通过本接口可以修改转化目标内的监测链接，
// 监测链接目前包含5个字段，action_track_url、track_url、video_play_track_url、video_play_done_track_url、video_play_effective_track_url，当传入值为空字符串时，会将对应的监测链接置为空，如果将5个链接全部传入空字符串，那么5个监测链接将全部置为空； 修改转化的监测链接，可以同时传多个链接地址修改，5个监测链接中，需要至少有一个字段传入，全部不传（字段没有传入置）系统会进行报错； 应用下载类的转化跟踪，不允许修改监控链接（展示，点击等5个均不允许传入修改）。
// 监测链接传输协议仅支持https类型
// 监测链接可以修改为空值，如 action_track_url传入空字符串，会将监测链接置空
func TrackURLUpdateStatus(ctx context.Context, clt *core.SDKClient, accessToken string, req *adconvert.TrackURLUpdateRequest) error {
	return clt.Post(ctx, "2/tools/ad_convert/track_url/update/", req, nil, accessToken)
}
