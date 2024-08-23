package v3

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	v3 "github.com/bububa/oceanengine/marketing-api/model/tools/adraise/v3"
)

// Set 设置一键起量: 用来启动或关停一键起量服务
func Set(clt *core.SDKClient, accessToken string, req *v3.SetRequest) error {
	return clt.PostAPI("v3.0/tools/promotion_raise/set/", req, nil, accessToken)
}
