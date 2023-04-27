package v3

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/adraise/v3"
)

// Stop 关停一键起量
func Stop(clt *core.SDKClient, accessToken string, req *v3.StopRequest) error {
	return clt.Post("v3.0/tools/promotion_raise/stop/", req, nil, accessToken)
}
