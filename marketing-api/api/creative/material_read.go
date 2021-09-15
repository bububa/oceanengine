package creative

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/creative"
)

// MaterialRead 创意素材信息
func MaterialRead(clt *core.SDKClient, accessToken string, req *creative.MaterialReadRequest) ([]creative.Material, error) {
	var resp creative.MaterialReadResponse
	err := clt.Get("2/creative/material/read/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}
