package clue

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/clue"
)

// FormGet 建站工具——查询已有表单列表
// 通过此接口，用户可以根据广告主ID获取广告主已有表单列表，可用于建站。
// 列表栏包括广告主ID、实例ID、实例名称、版本信息、是否已删除、是否未分层表单、表单类型、表单内是否包含电话元素、创建时间等信息。
func FormGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *clue.FormGetRequest) (*clue.FormGetResponseData, error) {
	var resp clue.FormGetResponse
	if err := clt.Get(ctx, "2/tools/clue/form/get/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
