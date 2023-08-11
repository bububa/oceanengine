package enterprise

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/enterprise"
)

// GetOverview 获取企业号基础数据
func GetOverview(clt *core.SDKClient, accessToken string, req *enterprise.GetDataRequest) (*enterprise.GetOverviewResponse, error) {
	var resp enterprise.GetOverviewResponse
	err := clt.Get("v1.0/enterprise/overview/data/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOverview 获取企业号流量来源数据
func GetFlowCategory(clt *core.SDKClient, accessToken string, req *enterprise.GetDataRequest) (*enterprise.GetFlowCategoryResponse, error) {
	var resp enterprise.GetFlowCategoryResponse
	err := clt.Get("v1.0/enterprise/flow/category/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVideoInfo 获取企业号视频分析数据
func GetVideoInfo(clt *core.SDKClient, accessToken string, req *enterprise.GetDataRequest) (*enterprise.GetVideoInfoDataResponse, error) {
	var resp enterprise.GetVideoInfoDataResponse
	err := clt.Get("v1.0/enterprise/video/info/get/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
