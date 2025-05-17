package star

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/star"
)

// ReportCustomDataTopicDailyReport 获取投后每日趋势数据（短视频）
func ReportCustomDataTopicDailyReport(ctx context.Context, clt *core.SDKClient, accessToken string, req *star.ReportCustomDataTopicDailyReportRequest) ([]star.ReportCustomDataTopicDailyReport, error) {
	var resp star.ReportCustomDataTopicDailyReportResponse
	if err := clt.GetAPI(ctx, "2/star/report/custom_data_topic_daily_report/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	if resp.Data == nil {
		return nil, nil
	}
	return resp.Data.Stats, nil
}
