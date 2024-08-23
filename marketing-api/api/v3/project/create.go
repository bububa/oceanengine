package project

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/project"
)

// Create 创建项目
func Create(clt *core.SDKClient, accessToken string, req *project.CreateRequest) (uint64, error) {
	var resp project.CreateResponse
	if err := clt.PostAPI("v3.0/project/create/", req, &resp, accessToken); err != nil {
		return 0, err
	}
	return resp.Data.ProjectID, nil
}
