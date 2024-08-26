package form

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/clue/form"
)

// Detail 获取表单详情
// 本接口用于获取青鸟线索通表单详情，可基于instance_id查询（instance_id唯一表示一个表单，可通过创建表单返回值中获取，也可从获取表单列表中获得），可获得表单配置的详细信息，不包含线索统计数据。
func Detail(ctx context.Context, clt *core.SDKClient, accessToken string, req *form.DetailRequest) (*form.Form, error) {
	var resp form.DetailResponse
	if err := clt.Get(ctx, "2/clue/form/detail/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data.Form, nil
}
