package enterprise

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/enterprise"
)

// ListMedia 企业号用户素材列表接口
func ListMedia(clt *core.SDKClient, accessToken string, req *enterprise.ListMediaRequest) (*enterprise.ListMediaResponse, error) {
	var resp enterprise.ListMediaResponse
	err := clt.Get("v1.0/enterprise/media/list/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
