package report

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/report"
)

// LiveGet 获取今日直播数据
// 支持获取直播整体数据
// 基于全部流量的情况(包括自然+广告),分析抖音号及直播间在竞价推广中的整体数据表现。直客账户支持查看官方抖音号下同主体全部广告账户的数据表现。
func LiveGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.LiveGetRequest) (*report.LiveStat, error) {
	var resp report.LiveGetResponse
	if err := clt.Get(ctx, "v1.0/qianchuan/report/live/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
