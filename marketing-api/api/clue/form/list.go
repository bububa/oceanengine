package form

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/clue/form"
)

// List 获取表单列表
// 本接口用于获取青鸟线索通表单列表，可根据时间和表单实例id等条件进行过滤，本接口仅获取表单等基本信息，需要详细信息需调用【获取表单详情】接口。
func List(ctx context.Context, clt *core.SDKClient, accessToken string, req *form.ListRequest) (*form.ListResponseData, error) {
	var resp form.ListResponse
	if err := clt.Get(ctx, "2/clue/form/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
