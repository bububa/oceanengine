package star

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/star"
)

// ReportCustomDataTopicConfig 获取投后数据主题累计数据
func ReportCustomDataTopicConfig(ctx context.Context, clt *core.SDKClient, accessToken string, req *star.ReportDataTopicConfigRequest) (*star.ReportCustomDataTopicConfigResult, error) {
	var resp star.ReportCustomDataTopicConfigResponse
	if err := clt.GetAPI(ctx, "2/star/report/custom_data_topic_config/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
