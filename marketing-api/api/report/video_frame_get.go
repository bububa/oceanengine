package report

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/report"
)

// VideoFrameGet 视频互动流失数据
// 此接口用于获取：视频互动流失数据，统计用户点击详情按钮或者左滑进入落地页的点击类行为发生在视频播放的第几秒。
func VideoFrameGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.VideoFrameRequest) ([]report.VideoFrameResponseDataList, error) {
	var resp report.VideoFrameResponse
	err := clt.Get(ctx, "2/report/video/frame/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	if resp.Data == nil {
		return nil, nil
	}
	return resp.Data.List, nil
}
