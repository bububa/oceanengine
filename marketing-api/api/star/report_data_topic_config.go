package star

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/star"
)

// ReportDataTopicConfig 获取任务下累计可查询的数据指标
func ReportDataTopicConfig(ctx context.Context, clt *core.SDKClient, accessToken string, req *star.ReportDataTopicConfigRequest) ([]star.DataTopicConfig, error) {
	var resp star.ReportDataTopicConfigResponse
	err := clt.GetAPI(ctx, "2/star/report/data_topic_config/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	if resp.Data == nil {
		return nil, nil
	}
	return resp.Data.Stat, nil
}
