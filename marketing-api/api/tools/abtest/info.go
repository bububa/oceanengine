package abtest

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/abtest"
)

// Info 获取实验详情及报告
// 该接口用于获取AB实验的详细信息与实验数据报告。
func Info(ctx context.Context, clt *core.SDKClient, accessToken string, req *abtest.InfoRequest) (*abtest.AbTest, error) {
	var resp abtest.InfoResponse
	if err := clt.Get(ctx, "2/tools/ab_test_info/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
