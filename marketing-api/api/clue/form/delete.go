package form

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/clue/form"
)

// Delete 删除表单
// 本接口可进行表单的删除，删除后获取表单列表时默认不获取，但是仍能通过表单实例id获取表单详情，is_del状态变为1。 注：对于当日产生了线索数据的表单不支持删除，会报错：message：has clues today
func Delete(ctx context.Context, clt *core.SDKClient, accessToken string, req *form.DeleteRequest) (bool, error) {
	var resp form.DeleteResponse
	if err := clt.Post(ctx, "2/clue/form/delete/", req, &resp, accessToken); err != nil {
		return false, err
	}
	return resp.Data.Success, nil
}
