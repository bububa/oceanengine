package smartphone

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/clue/smartphone"
)

// Get 获取智能电话列表
// 本接口用户获取青鸟线索通智能电话列表，可根据智能电话id，智能电话创建时间，删除标记进行过滤。
func Get(ctx context.Context, clt *core.SDKClient, accessToken string, req *smartphone.GetRequest) (*smartphone.GetResponseData, error) {
	var resp smartphone.GetResponse
	if err := clt.Get(ctx, "2/clue/smartphone/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
