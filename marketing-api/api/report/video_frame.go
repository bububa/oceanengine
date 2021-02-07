package report

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/report"
)

func VideoFrame(clt *core.SDKClient, accessToken string, req *report.VideoFrameRequest) (*report.VideoFrameResponse, error) {
	var resp report.VideoFrameResponse
	err := clt.Get("2/report/video/frame/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
