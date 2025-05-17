package customercenter

import (
	"context"

	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/customercenter"
)

// AdvertiserTransferableList 获取可转账户列表（客户中心&广告主）
// 本接口用于查询管家或者广告主的可转账列表注意：暂不支持代理商角色账号获取报错：账户角色不能为一代或者二代代理商
func AdvertiserTransferableList(ctx context.Context, clt *core.SDKClient, accessToken string, req *customercenter.AdvertiserTransferableListRequest) (*customercenter.AdvertiserTransferableListData, error) {
	var resp customercenter.AdvertiserTransferableListResponse
	if err := clt.Get(ctx, "2/customer_center/advertiser/transferable/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
