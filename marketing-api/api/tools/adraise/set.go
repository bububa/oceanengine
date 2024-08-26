package adraise

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model/tools/adraise"
)

// Set 设置一键起量: 用来启动或关停一键起量服务
func Set(ctx context.Context, clt *core.SDKClient, accessToken string, req *adraise.SetRequest) (enum.AdRaiseStatus, error) {
	var resp adraise.SetResponse
	err := clt.Post(ctx, "2/tools/ad_raise/set/", req, &resp, accessToken)
	if err != nil {
		return enum.AdRaiseStatus_UNKNOWN, err
	}
	return resp.Data.Status, nil
}
