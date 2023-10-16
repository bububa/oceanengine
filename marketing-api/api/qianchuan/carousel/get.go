package carousel

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/carousel"
)

// Get 获取千川素材库图文
func Get(clt *core.SDKClient, req *carousel.GetRequest, accessToken string) (*carousel.GetResult, error) {
	var resp carousel.GetResponse
	if err := clt.Get("v1.0/qianchuan/carousel/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
