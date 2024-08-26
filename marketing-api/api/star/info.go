package star

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/star"
)

// Info 获取星图账户信息
// 获取星图账户全量信息
func Info(ctx context.Context, clt *core.SDKClient, accessToken string, req *star.InfoRequest) ([]star.Info, error) {
	var resp star.InfoResponse
	if err := clt.Get(ctx, "2/star/info/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.InfoList, nil
}
