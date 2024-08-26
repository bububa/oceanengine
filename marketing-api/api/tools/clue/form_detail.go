package clue

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/clue"
)

// FormDetail 建站工具——查询表单详情
// 通过此接口，用户可以根据表单id获取已有表单的信息详情。
// 详情包括广告主ID、实例ID、实例名称、表单标题、提交按钮文案、副标题、表单元素、元素标签、是否允许为空、表单元素类型、是否可重复等信息。
func FormDetail(ctx context.Context, clt *core.SDKClient, accessToken string, req *clue.FormDetailRequest) (*clue.FormDetail, error) {
	var resp clue.FormDetailResponse
	if err := clt.Get(ctx, "2/tools/clue/form/detail/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
