package site

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/tools/site"
)

// UpdateStatus 更改橙子建站站点状态
// 通过此接口，用户可以更改橙子建站站点状态

// 更改橙子建站站点状态前，可在获取橙子建站站点详细信息的返回参数中获取当前站点状态status ; 站点状态的枚举值详见 【站点状态】，状态不同，可对应的操作会做相应限制。如DELETED 和AUDIT_BANNED状态下不可再上线；EDIT状态下不可直接使用。

// 新建的站点同样需要发布后才可生效投入使用！
// 恢复删除站点后，需要再发布才可生效！
func UpdateStatus(ctx context.Context, clt *core.SDKClient, accessToken string, req *site.UpdateStatusRequest) (*site.UpdateStatusResponseData, error) {
	var resp site.UpdateStatusResponse
	if err := clt.Post(ctx, "2/tools/site/update_status/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
