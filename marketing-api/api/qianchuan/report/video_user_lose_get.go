package report

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/report"
)

// VideoUserLoseGet 视频互动流失数据
// 此接口用于获取：视频互动流失数据，统计用户点击详情按钮或者左滑进入落地页的点击类行为发生在视频播放的第几秒。
// 互动行为越多的视频进度，说明这部分视频更吸引人；
// 分析披露的数据不包含重复播放后的互动行为，仅记录首轮播放时用户的点击和点赞、评论等行为所在的位置；
// 流失分析是指看过广告但是未点击的用户直接滑走了该视频，流失分析是统计这种流失行为发生在第几秒，即展示未点击的人都看到了第几秒。
// 不支持获取当天数据
// 数据更新频率:
// - 第二天9点可以获取前一天稳定的消耗数据；
// - 一般历史数据都不会变，除了数据有问题有校对的情况会更新历史数据；
func VideoUserLoseGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.GetRequest) (*report.GetResponseData, error) {
	var resp report.GetResponse
	err := clt.Get(ctx, "v1.0/qianchuan/report/video_user_lose/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
