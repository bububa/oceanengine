package customaudience

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/dmp/customaudience"
)

// Copy 推送dmp人群包到云图账户
// 通过此接口，将DMP人群包复制一个，并推送至对应的云图账号内，可在人群模块自定义分析查看和应用（该人群不支持再次推送）
func Copy(ctx context.Context, clt *core.SDKClient, accessToken string, req *customaudience.CopyRequest) error {
	return clt.Post(ctx, "2/dmp/custom_audience/copy/", req, nil, accessToken)
}
