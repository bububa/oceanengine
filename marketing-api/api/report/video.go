package report

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/report"
)

// 视频素材报表
// 此接口用于获取：视频素材数据，包括消耗、点击、展示等指标以及视频素材特有的指标，具体可以参考应答参数指标说明。
// 可以通过消耗、转化数、成本的多少进行过滤；
// 素材数据：不支持获取当日数据；
func Video(clt *core.SDKClient, accessToken string, req *report.IntegratedRequest) (*report.IntegratedResponse, error) {
	var resp report.IntegratedResponse
	err := clt.Get("2/report/video/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
